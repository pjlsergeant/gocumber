use strict;
use warnings;

use Path::Class qw/file/;
use Test::BDD::Cucumber::Parser;

use Data::Printer;

my $package_name = $ARGV[0];

my $feature = Test::BDD::Cucumber::Parser->parse_file( $ARGV[1] );

my $data;

$data .= sprintf(
    'f = Feature{
    Name: %s,
    StartsAt: %s,
    ConditionsOfSatisfaction: %s,
    Tags: %s,
    Language: %s,
    Background: %s,
    Scenario: %s,
}
',
    quote( $feature->name ),
    document_location( $feature->name_line ),
    satisfaction( $feature->satisfaction ),
    tags( $feature->tags ),
    quote( $feature->language ),
    ( $feature->background ? '&' . scenario( $feature->background ) : 'nil' ),
    (
        "[]Scenario{"
          . (
            join ', ',
            map    { scenario($_) }
              grep { !$_->background } @{ $feature->scenarios }
          )
          . "}"
    )
);

my $template = join '', (<DATA>);
$template =~ s/%PACKAGE_NAME%/$package_name/;
$template =~ s/%FEATURE%/$data/;

print $template;

sub scenario {
    my $s = shift;
    return sprintf(
        'Scenario{
        Name: %s,
        StartsAt: %s,
        Tags: %s,
        Background: %s,
        Steps: %s,
        TableData: %s,
        DocString: %s,
}',
        quote( $s->name ),
        document_location( $s->line ),
        tags( $s->tags ),
        ( $s->background ? 'true' : 'false' ),
        ( "[]Step{" . ( join ', ', map { step($_) } @{ $s->steps } ) . "}" ),
        table_data( $s->data, [] ),
        docstring( $s->data )
    );
}

sub step {
    my $s = shift;
    return sprintf(
        'Step{
        Text: %s,
        StartsAt: %s,
        Verb: %s,
        OriginalVerb: %s,
        TableData: %s,
        DocString: %s,
    }',
        quote( $s->text ),
        document_location( $s->line ),
        quote( $s->verb ),
        quote( $s->verb_original ),
        table_data( $s->data, $s->columns ),
        docstring( $s->data )
    );
}

sub table_data {
    my ( $data, $columns ) = @_;
    return 'nil' unless ref $data;
    return 'nil' unless @$data;
    return sprintf(
        '&TableData{
        Columns: %s,
        Data: []map[string]string{ %s },
    }',
        string_array(@$columns),
        (
            join ',',
            map {
                my $h = $_;
                sprintf( '{%s}',
                    join ',',
                    map { quote($_) . ': ' . quote( $h->{$_} ) } keys %$h )
            } @$data
        )
    );
}

sub docstring {
    my $data = shift // '';
    return '""' if ref $data || length($data) == 0;
    $data =~ s/\n/\\n/g;
    return quote($data);
}

sub quote {
    my $text = shift // '';
    $text =~ s/"/\\"/g;
    return '"' . $text . '"';
}

sub document_location {
    my $line = shift;
    return sprintf(
        'DocumentLocation{%s, %s}',
        quote( $package_name . '.feature' ),
        $line->number
    );
}

sub satisfaction {
    string_array( map { $_->content } @{ $_[0] } );
}

sub tags {
    return string_array( @{ $_[0] } );
}

sub string_array {
    my $internal = join ',', map { quote($_) } @_;
    my $external = "[]string{$internal}";
    return $external;
}

__DATA__
package gocumber

func Feature_%PACKAGE_NAME%() (f Feature) {
    %FEATURE%

    return f
}
