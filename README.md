# geoip Service

This is a simple service layer depends on <https://github.com/oschwald/geoip2-golang>.

Technically, it's less efficient than directly using the library
(since it adds the overhead of gRPC communication, and need an extra conversion step).
But it provides a language-agnostic interface via gRPC,
making it easier to integrate with applications written in different programming languages.
