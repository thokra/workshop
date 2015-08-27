# slide1
Les sjæl

# Historietime
Startet hos Google sent 2007 av Robert Griesemer, Rob Pike og Ken Thompson

Fungerende compiler og stdbibliotek i 2008.
Open Source fra 10. november 2009

Versjon 1 lansert 28. mars 2012

Versjon 1.5 lansert 19. august 2015. Uten C, litt lengre byggetid, men bedre kjøretid.

# Hvorfor ett nytt språk?
Utviklet pga frustrasjon med eksisterende miljøer for systemprogrammering. Det hadde blitt
vanskelig og treigt å utvikle/compile.

Det var lenge siden det ble lansert ett nytt, moderne systemprogrammeringsspråk som er bygd for multiprossessering.

# Hva er Go?
Go er en spesifikasjon på ca 45 A4 sider.
- Java: 644 (670)
- C#: ~500
- Python: ~120

VS PHP:
Nøkkelord:
- PHP: 70+
- Go: 25

Verktøypakke:
- Formatering
- Linting
- Bygging
- Debugging
Nøkkelord:
- Kommer tilbake til

# Hello World

# Syntax
Alle Go filer må starte med en package erklæring. Main er spesiell, det er her programmet starter.

Det er flere måter å definere variabler. `:=` metoden fungerer kun inne i funksjoner.

Funksjoner starter med keywordet `func`, navn(argumenter) returverdier

Slice er en liste av verdier som utvider seg ettersom du legger til. Array er en liste med x antall plasser.

Det finnes kun èn løkke, `for`. Denne kan brukes likt som `while` og `for` i andre språk. Vi har også tilgang til `range`

# Hello World 2
Implementasjon av server. Routeren som er innebygget i Go matcher så langt du har skrevet i `HandleFunc`

# Structs
Structs er Go-s "klasser" og kan inneholde data og metoder.
Metodene kan ha både pointer og value receiver. Hva velger man? Usikker? Gå for pointer

# Interfaces
Bygger videre på Go-s objektorientering. Alle typer som har metodene som er definert i ett interface, implementerer det implisitt.

# Concurrency
`go` nøkkelordet kjører en funksjon i en egen `goroutine`
En goroutine er en lett tråd, som Go holder kontroll over, og kjører på flere prosesser hvis det er satt opp