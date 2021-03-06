// +build js

// font used for letters: source code pro

package main


//all the images used, saved as byte strings
func init() {
	TileImgs = map[string][]byte{}

	TileImgs["letter-0"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAV0lEQVQ4jWNgGHSAEVPo////CGlG
dAWMuJTi0YOiAU0PpghhOTRxJpzW4QCUaYD7D4/TsdiA05e4NOALRzQNcLMZGRlxaaN9KKEAYmKa
5LREcmodBcQAAJn8O/F+864QAAAAAElFTkSuQmCC
`)
	TileImgs["letter-1"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMUlEQVQ4jWNgGCLg////////xyrF
gqaO+jYwkWrWqIYhqoERzsKfkBgZGcm0YRQQAwCp2hIN7jKm7wAAAABJRU5ErkJggg==
`)
	TileImgs["letter-2"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUElEQVQ4jWNgGHSAEY3///9/FGlG
dAUofDTVWPUgOBDVyNJw/ciCTPgMw3APigas0vg0EAlI1oAP/P//H2u4jTzVWBICdnWwaKVqPIwC
GAAABrky9E6EMlQAAAAASUVORK5CYII=
`)
	TileImgs["letter-3"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAZklEQVQ4je2SwQ6AMAxCN+P//3I9
mNgOEOvNg9yWPUrJNsbnNOEcEcv1RGA5Ay1tO9AwkkdsfgFWGjo0JoBkJTG1chzrEmSI2/uim/XS
VnMeVmKl4e6ZteGk2eNKm4Ra+vVv/dXRAVmVMASCyAPdAAAAAElFTkSuQmCC
`)
	TileImgs["letter-4"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVUlEQVQ4je2SQQrAIAwEd/3/n7cH
QVtJYhQ8CJ1jmMEgAa5CkqRhWALbnNtBtUmmgmabj7greYxBsIwRTO1PkLF7kLQB8B1MVBIbvxSx
dhoe54OfDA+wXicIiFgWcgAAAABJRU5ErkJggg==
`)
	TileImgs["letter-5"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWklEQVQ4jWNgGHSAEZnz//9/7IoY
EcqYSLWBBb95BAAuJyEDajgJ2R5M55EcSggWRDWakXAjSAsJZJtJ9jS5GoiJAQhgRNOA1dP4Qgm7
qVg1YNVGQmiOAiQAACQJJxDX8UETAAAAAElFTkSuQmCC
`)
	TileImgs["letter-6"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAYklEQVQ4je1SOw5AIQzSF+9/5b7B
TxSqVl0cZKRAWlLnroNnSkQahW80aAA124KqrlN7EWk2Gu+pP2teRmDK2tJgn9qjGCAy8oXEG0DN
WD76wBCXgdO5CdMv6S31PNMaHhg/nrowCtkOm60AAAAASUVORK5CYII=
`)
	TileImgs["letter-7"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAARklEQVQ4jWNgGHSAEc76//8/PnWM
UJVM1HcDms0EbMDvTiyqSdaAKYjTSaSZjUcDdhvwGE9yPFAp4khzEv7wGSg/jAJUAACYNSD2YoRs
XwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-8"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAZElEQVQ4je1QOQ7AMAiDqP//Mh0S
tYkxKMfSod7AB8gin4P6lZm9tKJAI+kg6mxXTkcRjaM07EvyLkWBmR7pszD1MVRRHacM0atDS5CX
VSQbLfnrFNjSmsF/RTrI6SY6qfXHDG47T0H7Um9AGAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-9"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAZ0lEQVQ4jeVSQQ7AMAiqZv//Mjss
2SoYPOyyZNwUrGC61ucQVAModLCg1KRuZ57iVu+Kq7l3cnhPLKWn1SpvGFEGNLR23Ib+aF4REXQo
F7q9waFPGpMzAJDPNNwQuqXVpPutb/P8GCepbDMQH0KS6AAAAABJRU5ErkJggg==
`)
	TileImgs["letter-A"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXElEQVQ4je2RwQ7AIAxCZdn//zI7
LM6lJaS60xI52lIesbWfiCRJOTrktrklDN6pDQCqCc/V25NDHFIpwfNEQwCQVNNIZ37y/4DqXi8W
E3Ljrx3GGYP0nq4mbFldSrMtA7gBVywAAAAASUVORK5CYII=
`)
	TileImgs["letter-B"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVklEQVQ4jWNgGHSAEc76//8/dhWM
jMhcJoJGohmEbgOaeXDVcHECNqDpJ8pJpGnAdCftQwlFAqscmjhhG0h2EprDiPUDsRGH6SvCwcqA
6g2SPT0KiAEAha0tCvChWkMAAAAASUVORK5CYII=
`)
	TileImgs["letter-C"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXklEQVQ4je1SQQ4AEAxD/P/LcxCx
VRkHiYMet3YbbQjPIdKqiHRGNBwUaKrhNVmmbD0VRpAG3ABIs9W7AhfHgo7Z/6w27GjuvwEF9CqT
FNqgTtficZaIrxgez/uPEQW1dycNRNauhwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-D"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWklEQVQ4je1SQQoAIAhz/f/PdghE
VNQOQod2CnObsBE9B8iLmYNvwExWrudVrIOWlG09zBwAnFXtU5zkUROMyYDDPMFE0XWQKK6Tzghh
0kX5yPWvOMm39aODDW2SHhymX90sAAAAAElFTkSuQmCC
`)
	TileImgs["letter-E"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAOElEQVQ4jWNgGHSAEZnz//9/nOoY
oSqZSLWBBY9hWAHJNtBeA1GhhOyrERlKtAlWBkoS3yggBgAAMUUJJrlUpZ8AAAAASUVORK5CYII=
`)
	TileImgs["letter-F"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANElEQVQ4jWNgGHSAEZnz//9/nOoY
oSqZSLWBBY9hWAHJNtBeA1GhhOyrERlKg1DDKCAGAADE0wYn9o4BfQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-G"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAYElEQVQ4jWNgGHSAEVPo////KCoY
UdSga0BTjakHRQNcNbIKiCCaPVAJrMajASaCKrADIo0nxwYWPHYic+Gepp4NcCPRrGJClqZJsCI0
wC3Bbw9RaYkByUskp9ZRQAwAAHQ1LQu5bs34AAAAAElFTkSuQmCC
`)
	TileImgs["letter-H"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMklEQVQ4jWNgGMTg////////JyjO
RKq5w0EDCxofa0BR1QZGRkb8dg7CUBqEGkYBMQAAsgISErMyrJwAAAAASUVORK5CYII=
`)
	TileImgs["letter-I"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAALklEQVQ4jWNgGHSAEc76//8/PnWM
UJVM1LH3////uCwk2YZRDUNUw0AlvlGACgDYYQwVeZqOCwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-J"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAQUlEQVQ4jWNgGHSAEZnz//9/7IoY
EcqYqOyA////o1lLsg2jGmisAVe6wGcDWrxiNYJw4kNOeegaMPWgqR4FRAIAuD4bCGleye0AAAAA
SUVORK5CYII=
`)
	TileImgs["letter-K"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAYElEQVQ4je2RMQ7AMAgDTf//ZzpU
QgjseumQoZ6ixIeNAhynqFNmAoiI4Xju6+l6nzfcBthun7BLSkCtxAHl5kBVp5pAn01JuYNiCECr
S6C7aYj/B19pB/aQrxNoyC+rGyyoJyHwcT7qAAAAAElFTkSuQmCC
`)
	TileImgs["letter-L"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAALUlEQVQ4jWNgGNzg////////x6+G
iVRDRzWMasABGJE5eFIeIyMjmTaMAmIAAGP8CRaVjXzxAAAAAElFTkSuQmCC
`)
	TileImgs["letter-M"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWUlEQVQ4je2RsRIAIAhCs+v/f5mG
NkStobuGGMkXZK09LAAASr+Xt5DDgJ8wsxCgM6ms0kqjTAHIp4cAtfIlw0pRjgbWtNyBAPJdFR/n
NaSbhBwn3Ae+djQBsEshE8dQQO8AAAAASUVORK5CYII=
`)
	TileImgs["letter-N"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAATklEQVQ4jWNgGMTg////////JyjO
hCmN31wUDYyMjAQdgqIBYjx+S9CdRJoNcFfhsYRiGwhaQg0b8FtCJRvwWILPBqwRTz0n4bFkFBAE
AE8SIRFoIhqWAAAAAElFTkSuQmCC
`)
	TileImgs["letter-O"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXUlEQVQ4je1SsQ0AIAgD4/8v46BB
pGrUicFuQksrgSgcGEsi0tvsCbyiDiQjy8i2bRyRNsP0aWVpHwYxOOAX0cRHunMIKJjuqhZ1H00w
3c9RJGfixtPDLV1f68cJChMcJxodPcJCAAAAAElFTkSuQmCC
`)
	TileImgs["letter-P"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASElEQVQ4jWNgGHSAEc76//8/dhWM
jChcghrQtLHgNw/TICY8pmIFhDWg2UlYA2lOwgwJokIJ2VUk+4FwsKIBGgTr4NcwCogBABUgEiPV
tqjTAAAAAElFTkSuQmCC
`)
	TileImgs["letter-Q"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAa0lEQVQ4jc1TQQ6AMAgD4/+/XA8z
E1tBZ7JkvW20UBgzWw6uVwCusDPBM+qNFGS7smNYU2xFsn6MMhZ8BYCsAYoOV1hB0Nz36Z0CHd80
S1SE/NiPXeJlJI1u6zPqhx/TpBWV3ey9WKx/3xwcePI/Ao8j9xcAAAAASUVORK5CYII=
`)
	TileImgs["letter-R"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXklEQVQ4jWNgGHSAEc76//8/dhWM
jMhcJoJGohmEbgOaeXDVcHECNqDpJ8pJNNaAGXSEg5UB1SeEnYTmbxZc0lhDGZ8NEKWY7iQ5pvFp
wHQPYRswHUZsxOEJ9FGACQBX7iEeyp8VgwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-S"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAX0lEQVQ4je1SORIAIQgL+/8/Y+Os
CBF1bCxMKTmYIHAdJD6paseQjiMJlWqIwI5/C5dTZ6MEi2/KyLAS4jezArL3aa0jJY3KZDZzu6Um
WDmCT4i1Rpezv8Qttyp6AAAUo1o485e+kI8AAAAASUVORK5CYII=
`)
	TileImgs["letter-T"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAALUlEQVQ4jWNgGHSAEUL9//+fsFJG
RgYGBibq2Pv//39cdpJsw6iGUQ2jgAIAAGv1CRbyg0exAAAAAElFTkSuQmCC
`)
	TileImgs["letter-U"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANklEQVQ4jWNgGMTg////////JyjO
RKq5oxpGqga05AThMjIywkUY8ahGKMKlAVMPstJRQDwAADJNGAxgBBDNAAAAAElFTkSuQmCC
`)
	TileImgs["letter-V"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXElEQVQ4jWNgGKzg////////J0aW
iRizGBgYGBkZUTRA+HgsgQPCNpCjAe4eFA1YXYXpSBo4Cdk96BrQXIU10KjtJDT3YNEAV4ErErHb
gCfKSfYDdoA/8Y4CggAA+e0qCWwjtQMAAAAASUVORK5CYII=
`)
	TileImgs["letter-W"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAZklEQVQ4je2SOw7AMAhDH9z/znSI
hCp+TaUMHeqNYGOHBD4HAcwMEJGOtAiLo/m0nS0C6Dw7Q58pOwIz6xKecPDZpcmugy9GQ12WrcM9
g2tCsBgptPM1TmxpfvjJoVRuRXr12X4AFyN8IR4XoFjLAAAAAElFTkSuQmCC
`)
	TileImgs["letter-X"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWElEQVQ4jeVSuwoAIBDy+v9/tqGh
6K6saAhyFHwgAg+DJEnJJ+kCwMwCQWHDkNViYUlRycM81bq27Q8TYowmvlHJDy8EnbeYVdrPEiYh
1X7jfPJC2x/7Gxk2V0flHa2atwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-Y"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAQ0lEQVQ4jWNgGKzg////////J00W
lx40cSYqOwxThDIb0IzE6iuSbWDEaglCmhFdAT4bMFVj14BVHVE2DFsNo4AYAAAf0TLstDSWXgAA
AABJRU5ErkJggg==
`)
	TileImgs["letter-Z"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANUlEQVQ4jWNgGHSAEc76//8/PnWM
UJVM1HfD////8Vs+4Kqp5OkB9ShlfqCJ0wdD4hsFDAwA0xk71RqHAUgAAAAASUVORK5CYII=
`)
	TileImgs["letter-a"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVUlEQVQ4je2QsRIAEAxDxfn/X46N
0lAGm0yufbkmUvp6IPgRyYHAwGCDSk9/Ndqu5TAQSXu5rKCVvxxywmBpWUNHCsvl/dqHnC+ENfoF
HwbAxfd/GVXY8SQT1EK8RQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-arrow"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASUlEQVQ4jWNgGKHg////////J1Ix
I0QDgs/ISJoGgjpxasCljYAGTJ3EaoBrYyFeKYRBQAOmH3BqwBVK6BoIxgNCA0Glo2AwAQCYwBUj
NrCdMQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-asterisc"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAATklEQVQ4jWNgGK7g////uKSYSNWD
UwNhG/A4A1mKBVOCkZERjxGMeAyDqmBEUYOuAU0PmmoGMjyNxXi4DXiCAbt7sOphwi9NrD2jYNAD
AJfSLO2WOBYXAAAAAElFTkSuQmCC
`)
	TileImgs["letter-b"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUklEQVQ4je2RwQoAIAhDW/T/v7xu
ojaQLlHQO65NcrZ2DpIkS1vfnXthYEjVbw/AP0GaEj6jA+Yw0ZS8A4AwL/5HBEoeCKRy16637yBq
LZVPyQTMWRsh/QC8EQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-backslash"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASklEQVQ4jb2SMQ4AIAwCif//M25d
VJJSIvvlIC3wLyRJ5pnlVypJ26CZcSXTIJhEJciDhAx4z1CGK5OrZBqch+8ZTiY6ugztJaNsXZ41
4Q5NDtkAAAAASUVORK5CYII=
`)
	TileImgs["letter-boxe"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAALUlEQVQ4jWNgwAb+//////9/rFJM
WEXxgFENg0MDC64YpZoN2MFoWhrVQLoGAHeeDyBp6/G3AAAAAElFTkSuQmCC
`)
	TileImgs["letter-boxne"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAKUlEQVQ4jWNgGAUjAzD+//8fpxwj
I6YgE3Xs/f//Py6bSbZhVMNI0QAAUjkJFKC9N6EAAAAASUVORK5CYII=
`)
	TileImgs["letter-boxse"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAKklEQVQ4jWNgwAb+//////9/rFJM
WEXxgFENg0MDC64YpZoNo2AUUAsAADXMCRDEi5wNAAAAAElFTkSuQmCC
`)
	TileImgs["letter-c"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAARUlEQVQ4jWNgGAU0AIyYQv///0dR
wYiiBl0DmmpMbSxYVSObissIhv///+OUQwJMBFUMfg0owYonlOAiJMcDyTE9CogBAPxmHgkhag0h
AAAAAElFTkSuQmCC
`)
	TileImgs["letter-colon"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAKUlEQVQ4jWNgGAXkgv///////x+r
FCNW1QhpRnQFTDR30vAFo/Ew/AAAdWUd7xfXyb8AAAAASUVORK5CYII=
`)
	TileImgs["letter-comma"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAOUlEQVQ4jWNgGAWjYISB////////
H6sUI1bVCGlGdAVMWMzAUERAA3492DWQDPB4mmLVDKihRAUAAAJ9F/ke+NToAAAAAElFTkSuQmCC
`)
	TileImgs["letter-d"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAZElEQVQ4je1SOxaAMAxK8rz/lXFQ
Y4TWp0vrIBstkE9rNh8AACSNt/4PGhbidT4zc3cSXDipm86zQqprqkbwDNSDthQ1Sa8Vw95ha+Zm
S90K1UPfbo/uqVl37COap0mf7O2HYgX8gSQeOVIZ4gAAAABJRU5ErkJggg==
`)
	TileImgs["letter-dot"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAKElEQVQ4jWNgGAWjYGiD////////
H6sUI1bVCGlGdAVMNHfSKCAGAABHNg740fM+0wAAAABJRU5ErkJggg==
`)
	TileImgs["letter-dots"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMUlEQVQ4jWNgGAWjYNgCRgj1//9/
KJ+REc5FZsO5TJhmwFXAGcgAiwaISciMUUAqAABj+g8G/3WJvwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-dreaming"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAYUlEQVQ4je1SQQrAMAgzZf//cnYZ
UqMWStlhsJyqJia0Nfs+IDXJMIYSQi3shexhu6A7m9lY787moxt0W1Qwxy2jX8I+cihNggNJAPUl
dg6CnAp55g4z25vbL739l368ghu7TTwDVNiIYAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-e"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVklEQVQ4jWNgGAU0AIyYQv///0dR
wYiihhGPUqx6ECy4ajQjIeJwQSY8hmEVYUGTxuUqOEC3gSBAtwHTSZTagK4B0w////9HFiQqHhiQ
nEpyTI8CYgAAYc8eE5PHb5UAAAAASUVORK5CYII=
`)
	TileImgs["letter-equal"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAJUlEQVQ4jWNgGAU0AIxw1v////Gp
Y4SqZKKte8gBw8EPo2CIAgANWQYH3bjBwAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-exclamation"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAM0lEQVQ4jWNgGCLg////////xyrF
RKpZoxpoomGgAJ6kwYhVNUKaEV0BlTyNx0mjgBgAAO6MFQKPyEb8AAAAAElFTkSuQmCC
`)
	TileImgs["letter-f"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAQ0lEQVQ4jWNgoDVgxCP3//9/FKWM
jAwMDCzEKEUG2DUgG4lmBBMeDbg0k+YkLJbit4FYJxEAeGwj2YZRDTTRMAqIAQBpihIZqcW/PgAA
AABJRU5ErkJggg==
`)
	TileImgs["letter-fog"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAc0lEQVQ4jc1TQQrAMAgz+/+f3WHQ
itFYBoN5syaNsdbs6wAfufsuIwNyHtEdJxMiJ6UD+ihGztXROgIYB0AMCoxmnXhYt7RwDzQKKkIZ
m7Au01NSq7HcRw/ts5eOTXgYVmjUqRV48wplXebehv/w0snP4waqE1DxUsNqugAAAABJRU5ErkJg
gg==
`)
	TileImgs["letter-frontier"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAV0lEQVQ4je2RwQrAIAxDE9n//3J2
EB3WEnoU3LvVNDVV4GokVSSGI5KLPEqSi8FfMrsBPF7eRzTTnfIZ0jy7vxmtFMk8bidZ2nviDiFY
L5MRkoo//XMGL4jGMAPFTm3JAAAAAElFTkSuQmCC
`)
	TileImgs["letter-g"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAbElEQVQ4je1RQQ7AMAhSs/9/mR2W
OGcoNVl2GzcrKFSzHx/AWw1gSXU3sxiyE8dqUh1RX0L0KEK3yf5aZIZrTyuJwFhuZTLZAPY/NiJV
S5qdxnhoobkPJ+5QBz3usL1aF1BL7fFdaCqb+FQ4AfRHPAUIlersAAAAAElFTkSuQmCC
`)
	TileImgs["letter-gt"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAR0lEQVQ4jWNgGMTg////xChjxNTA
yMiIQzGqBuL1YHEbHudhNwmPVTitpprzmEjQTZ6TSPY0dtXEBisxBpOcNLAYPwoGAAAA570v86cV
AugAAAAASUVORK5CYII=
`)
	TileImgs["letter-h"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASUlEQVQ4jWNgoB/4////////CSpj
ItXcQaiBBasosu8ZGRkJ2IAWVmhcRqwScFPhgnARLDYguwHNPVg0YKogoIEgGJEaRgExAAAzTxUd
exIgrAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-hash"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAV0lEQVQ4jWNgGArg////yGxkLgMD
AxNB/YyMjKRpQAMka0BYh+ZWdHUwh5FsAzpACxZMa7HYALcdqyNp42nkqGDB7yS0WMMC8PuYgdKY
xh93ZNowCogBAK8vJwTnOpfAAAAAAElFTkSuQmCC
`)
	TileImgs["letter-hbar"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAHUlEQVQ4jWNgGAUjAzD+//+fJA1M
NHLIKBgFBAEACLQDAWnx+3YAAAAASUVORK5CYII=
`)
	TileImgs["letter-hit"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXklEQVQ4jdWQMQ7AMAgDQ9T/f5kO
XRA4jlGUoZ4YfIA9hiZ3/4bZcqtAW0cXLgDxH0kJaL/01E1mRoAZ3dW6D5AcFQAZ+NYM8AD4Aj+F
AdIYAKRadS2BVVcY2Hb1b72SASoLxAb2+wAAAABJRU5ErkJggg==
`)
	TileImgs["letter-hyphen"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIklEQVQ4jWNgGAVDEjDCWf///8en
jhGqkom27hkFo4CKAACI8QME8uqXXgAAAABJRU5ErkJggg==
`)
	TileImgs["letter-i"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAPElEQVQ4jWNgoBb4////////McWZ
cKnGZRB2DYyMjOS5a0AAwq1YPYrpGeyeJgeQFg94wKgGmmgYBcQAAHArEgyeM0JuAAAAAElFTkSu
QmCC
`)
	TileImgs["letter-interrogation"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAATUlEQVQ4jWNgoDVgxBT6//8/QpoR
XQEjLqW49GDRAFcB149pD07w//9/NGuZiNVKnvEDpZoBR/hCAFU9TR2ALy1hjWCS/UB7DaOAGAAA
H88p9IuKJvgAAAAASUVORK5CYII=
`)
	TileImgs["letter-j"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASklEQVQ4jWNgoBb4////////McWZ
cKnGZRB2DYyMjOS5a0AAwq1YPYrpGeyeJgeQFg94wKiGoa0BLRXgyXGMeFRgzXc4kzfVcikAU6Ae
DvI9I/kAAAAASUVORK5CYII=
`)
	TileImgs["letter-k"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWklEQVQ4je2SMQ4AIAgDxfj/L9fB
hBhAi4txsPMdWLWUewEAgGL1dO6DQqOE3oSIcGHQA+VH8vROCOmlsKJjQVuGDy+e09nZDoYwe6ww
09kO3p+X8K/hnR+aDgQAKhx0pOT0AAAAAElFTkSuQmCC
`)
	TileImgs["letter-kill"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUUlEQVQ4je2QOQ4AIAgEd43///Ja
EI8g2FmYOI1xQA6BzwVoh6SpeJJ0gREOJYDq7mFV85JIFt+xP1uzV8o+bmaMesgI+6Qfki0dD5pV
+bxLAy0ULQ4kMZu1AAAAAElFTkSuQmCC
`)
	TileImgs["letter-l"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAOElEQVQ4jWNgoDVghLP+//+PLsfI
yIABmKhj7////zEtJNOGUQ3DWQNyMmEhUh0BG7Am7FFAPAAAG1kPF1Y3zLEAAAAASUVORK5CYII=
`)
	TileImgs["letter-lbrace"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVElEQVQ4jWNgIBEw4pL4//8/ijpG
qEomYlQjAxZ8tjNisR+7DXjAqAYIwBMJDGgxjawUaySQ4yQsxsDtITamcTmGTCcNBw3YAwQzsglk
UfwhS2MAAOcSEiyKk3AcAAAAAElFTkSuQmCC
`)
	TileImgs["letter-lbracket"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMElEQVQ4jWNgIBEwYgr9//8fizpG
qEomUm1gwWk1IxbLybFhVMOohmGsgfZZlPYAAIUdBi6L64lcAAAAAElFTkSuQmCC
`)
	TileImgs["letter-longhyphen"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIklEQVQ4jWNgGAWjACtgZGBg+P//
P7GqGRmZaOmaUTCiAABzmAMEIf36fwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-lparen"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAT0lEQVQ4jcVSQQ4AIAjC/v9nO7aV
MTVdHBVQnEAtVHWriIctwmiLfdoDGN7lcvZMcGsZK3HvcIZ+gYFMhtiVOL6ERi5D7JfQ/t58SLXm
FRPxJjXlQ5S8QQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-lquotes"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANElEQVQ4jWNgoDf4//8/fhEmMvSQ
ppQJjxxOg0nSwIgpzcjIiEuEgQqeRpPDFBkFo2DoAACmmjXYy1xXEQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-m"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAP0lEQVQ4je2MMQ4AIAwCD///5zo0
6YAmOujWWwo0AM0HlCciAEkpUlde1guHbQkYa1pjq/WC/cxuCjd0oXnFBBkkDCfYa8CbAAAAAElF
TkSuQmCC
`)
	TileImgs["letter-magic"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAX0lEQVQ4je1SSQoAIQwz/v/PmcNA
O2osLXgRJidpli7Y2iUgSVJSvZqFNXtWYNAgkEpbn9QAjPu+DcMOK21FSywv7QYZHxnipYWhPNKu
yVv00yVHEoadLXWMk5/vRwYPsHonFL9lilYAAAAASUVORK5CYII=
`)
	TileImgs["letter-music1"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAT0lEQVQ4jWNgGDDw//9/rOJMpOph
JEY1IyNCGT4b4OqQNePTgGY2URow9RDWQLINlGrACciJOOpoYCHGGdhtIEY1PidhxjE+DbhUjwIi
AQC9MRUbP+PYIwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-music2"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAVUlEQVQ4jWNgIBEwElTx//9/hGpG
Rhb8KjABE6lOIlkDFichA0ZGqCfh7sSuAa6OCk4akRoIxANmMmHBJYEpCIkcJlyqcQFi/QCPe/Sk
gidRjALiAQAafhgllrVjzwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-n"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAARklEQVQ4je2OOwoAMAhDk97/znYp
QVSwS6f6JpV8BIYHUJOZASCp+ShIb1g5w6vzGhtCqo66FA3+h/BPYciKxtDypWG4YQPMpxIYUe3D
0QAAAABJRU5ErkJggg==
`)
	TileImgs["letter-o"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAT0lEQVQ4jWNgGAU0AIxo/P///6NI
M6IrQOGjqcaqDcGCq0aWxhRkwmMYVicxIZuEKY0J0G2gmQaIY3CFEj4b0PRgGkFBPGDVQ0y4jQJM
AAB4jR4S7eFipgAAAABJRU5ErkJggg==
`)
	TileImgs["letter-p"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAU0lEQVQ4jWNgGAU0AIxw1v///xkY
GBgZGeFsqApGRnwasJuKpAe7BrgKuCBchAnTMBTzUN2DRQNBMAQ0oAUuZliTHA9YghW/CAt+8zDB
EAhW6gMAcuoYKyywTUoAAAAASUVORK5CYII=
`)
	TileImgs["letter-percent"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAbklEQVQ4je1RuxLAMAiSXP//l+nW
GpRcMnSrk0/gMOLrwJORfLtAtxwRMep2LXNnTHRAi13vVwwkczkxOOxMO2RWExEJp9IZdWyrHTgI
e+BEqktiovtMs+0+o7au8PJBu92yXVluK1I6x7b+sRM3utVW7yRjBK8AAAAASUVORK5CYII=
`)
	TileImgs["letter-pipe"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIUlEQVQ4jWNgoAr4////////sUox
kWrWqIZRDaMaRrgGACQhBiczCAYQAAAAAElFTkSuQmCC
`)
	TileImgs["letter-@"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAe0lEQVQ4jc2TQRKAMAgDi9P/f7ke
6iAECD05coSEBapj/C4kptZaTiFOgwZQR48zqNoqdhI4TyFtD9grVZBw6Mi1vXd1xkKVYSPZ7QE7
K3VyFkIg0RtgmWQkESFv0hOqZco7MsKJB0dqPa9BZwUP/h6HBG2XXyB+cx/GDXe0RQgQbFivAAAA
AElFTkSuQmCC
`)
	TileImgs["letter-door"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMElEQVQ4jWNgGAVQ8P///////2OV
YiLVLNprYISzcDkaqo6RkUwbsIOhFUqjYIgCACGbDwS8lF+4AAAAAElFTkSuQmCC
`)
	TileImgs["letter-portal"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAWElEQVQ4je2SMQ4AIAgDxf//uQ4O
Ji1iVBYTGaG1R6SURwoAAHdkrnqMjQU1h0ejpgkkUs82kmPQRacGAuhOal4g9ZdinnWCUp1+XHA8
lGDaivEybumXVANG7Sz8egYphAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-q"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAX0lEQVQ4je1RQQ7AIAizxv9/uTuQ
6FZQ4wFP68kCLaSW8iMBEE7y0wasAiAQyLQo7dH8dO+FFnXmFNIhkEMX0A1pAjtmkdJ0w1tD0lsc
/0MNq5363JrwbbLXYk0UHOMB6oshKB0gbnMAAAAASUVORK5CYII=
`)
	TileImgs["letter-quote"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIElEQVQ4jWNgGBjw////////Y5Vi
ItWsUQ2jYBQMLwAAzosGCaECG94AAAAASUVORK5CYII=
`)
	TileImgs["letter-quotes"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIElEQVQ4jWNgoCv4////////8Ysw
kWroqIZRMAqGFwAAaFEMAyr1H9UAAAAASUVORK5CYII=
`)
	TileImgs["letter-r"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAP0lEQVQ4jWNgGAU0AIzInP///zMw
MDAyMiJzoepggixYjUFWStgGTFORAXYbsCqFACaSVGPXgB+MSA2jgBgAAKdKDBzeBS9nAAAAAElF
TkSuQmCC
`)
	TileImgs["letter-rbrace"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUElEQVQ4je2TwQoAIAhDZ/T/v7xu
UeYoiW69k6BTcQgksR6RnBJmSzEAFNXJ6TeQVAI5QfEFjvCysZ2uenQ9vZJs/52+EdQeHb6YnKB+
+j0NMF4bFZFcrnoAAAAASUVORK5CYII=
`)
	TileImgs["letter-rbracket"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANklEQVQ4jWNgIBEwwln////HIs3I
iCbCRKoNOMH///+x2kmyDaMaRjUMYw0scBbW7EIFG2gPAAKXDCGXzNPFAAAAAElFTkSuQmCC
`)
	TileImgs["letter-rock"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAQElEQVQ4jWNgGIGAEVPo////CGlG
dAWMuJSiKELSxkRQNRpgIqwE1SyiNJBsA+01YA8lzCDHrhlTCH/EjYIhCgDqYQ8SXsVRJwAAAABJ
RU5ErkJggg==
`)
	TileImgs["letter-rparen"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAPUlEQVQ4jWNgoAT8//+foBomMvSg
axhoPeh+IB+QYwNWPSQ7aURqwAlIiwc8aYkaTiIttdJSNQMZ+ZMmAACTqjjeE0bFngAAAABJRU5E
rkJggg==
`)
	TileImgs["letter-rquotes"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAN0lEQVQ4jWNgoCv4////////8Ysw
IsshRBkZsYowMDAw0dxVlNlAjKfJVc2AGia4REbBKBhSAAA0wjvQeMkoywAAAABJRU5ErkJggg==
`)
	TileImgs["letter-s"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUElEQVQ4je1PRw4AIAgD4v+/jBdj
mI4Yb/TWdFAACh+AhjOzktEaFDfuMNa8W8pZxdBW8naSXx8EwhkmGdf48IxRVpwNI0mens7uXD9d
OEEH4O4nBElOWW4AAAAASUVORK5CYII=
`)
	TileImgs["letter-semicolon"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAOElEQVQ4jWNgGAXkgv///////x+r
FCNW1QhpRnQFTDR30iiAAxpHHKYiAhrw6xmQpEFyQqJysgMAMKwm8Ngar/8AAAAASUVORK5CYII=
`)
	TileImgs["letter-tree"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAUElEQVQ4je1SQQ4AIAiS/v9nulVT
c+XyFjcVUZwi1YBNkZxlaIKOV7bb025Xqm/ImrZeHW0gu9KJ/Bjy7kqOXUASpn2Q3EnUPx8CajDn
I0YHe/QhFvreLOwAAAAASUVORK5CYII=
`)
	TileImgs["letter-slash"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAATUlEQVQ4jWNgGHjw////////k6+a
ifruQRPBZwMJTsejAacNuIynnqdJswGPd6nhJPxpgUo24JFFt4Fg7FLsJNJsICa1UeAk0rIiDQEA
yzM14X/tLloAAAAASUVORK5CYII=
`)
	TileImgs["letter-space"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAEklEQVQ4jWNgGAWjYBSMgsEFAASY
AAHKGPyHAAAAAElFTkSuQmCC
`)
	TileImgs["letter-stone"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAIklEQVQ4jWNgGAWjYBQMYsAIof7/
/09AHSNUJRNt3UMXAADN0AMEvr/DMAAAAABJRU5ErkJggg==
`)
	TileImgs["letter-sun"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAaklEQVQ4je1RywrAMAgzo///y+5g
WW18QNlpsJyqibEYkY9AVVU1pa5TrzVQWRIFagHYaFcCWAP9EpMaRk9HC/AXAVIQe3yl0qxqvsgh
9Ys7B3GUQ7nhUct+KHv3FnPei3yZBBcRw/nR4wZrPkT7ECoX6wAAAABJRU5ErkJggg==
`)
	TileImgs["letter-t"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAQ0lEQVQ4jWNgGArg////////xyXL
RKpxtNfACGfhcTcDAwMjIyOZNmABQy6UBokG5EBjIUYRMmDEKopVAzyyRwFBAADJZRgJl37VyQAA
AABJRU5ErkJggg==
`)
	TileImgs["letter-tick"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAN0lEQVQ4jWNgGAWDATBiCv3//x8h
zYiuAJ2PrBqrHiZSnYSiAdN4AhowXUxAAzGA5FAaBUMUAABNRgwMbG6kRQAAAABJRU5ErkJggg==
`)
	TileImgs["letter-tilde"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAANElEQVQ4jWNgGAWDDvyHAUxxOJsJ
j05M1QwMDIxoxjAyMmJRxIhQxoRVAlkFMnsUjAIqAgAtMBr6apZkYwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-times"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAPklEQVQ4jWNgGAUkgv///xMUZEKT
QJPGagS6eXBFyGw4YMTvBkZGLAoI2ENtGwj6gYBL8OnBJUc4ZEfBUAAA5vJHw8EB1xcAAAAASUVO
RK5CYII=
`)
	TileImgs["letter-u"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAOElEQVQ4jWNgGAU0Bf///////z9B
cSZSzR3VgAkwAxqfBqzRwoJfESMjI5oIE7IcmmZMkVFAJAAAthQYCKM54loAAAAASUVORK5CYII=
`)
	TileImgs["letter-v"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAARUlEQVQ4jWNgGAW0A////////z9p
srj0oIkzUdlhmCKU2QAxEg8Xuw1wRVjDgICTGBkZ0UXwuwpTA8WehluCJ+JHAUEAAMxjONnsXb+d
AAAAAElFTkSuQmCC
`)
	TileImgs["letter-vbar"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAJElEQVQ4jWNgwAb+//////9/rFJM
WEXxgFENoxpGNYxqIA0AAFHYBil6UycHAAAAAElFTkSuQmCC
`)
	TileImgs["letter-w"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAATUlEQVQ4je2PMQ4AIAgDhfj/L9fN
KAWMiW7cBvQItFZ8AgCAu4wReMUsNXLc9Cbw2F3hCDm+ICJRR48JQ18LPpo7L36IjrFC8ka+omAG
g0kqA48OK3IAAAAASUVORK5CYII=
`)
	TileImgs["letter-x"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAASUlEQVQ4je1QOQ4AIAij/v/PdWAx
SComutGRXoBZ4ytIkqzPc06pd0WqHnpPAHGSlgjDebHXDXdHO61T1ZfcGUqwZqfxldpGwAR3ykfX
q5fnOwAAAABJRU5ErkJggg==
`)
	TileImgs["letter-y"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAXElEQVQ4je1SQQ4AIAjS/v9nOrRV
K7Q565a3FBQxkR/vAgCAWNXiLPlyWdieyU1oLZ0nn9BB1IOVoKq+BL5D773zzaWdO3K0RQjbmr6D
nD6izriRtc0lko6niEUFvNpE231dSPwAAAAASUVORK5CYII=
`)
	TileImgs["letter-z"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAIAAAB8wupbAAAAMElEQVQ4jWNgGAU0AIxw1v////Gp
Y4SqZKKyA/7//4/f5lHVWMGAR9woYGBgYGAAAEc+L93Qbp4XAAAAAElFTkSuQmCC
`)


//fallback tile
TileImgs["map-notile"] = []byte(`iVBORw0KGgoAAAANSUhEUgAAABAAAAAYCAMAAADEfo0+AAADAFBMVEUAAAD///8A/wD/+wCC/wAA
/wAA/30A//8Agv8AAP95AP/PAP//ANf/AIL/AAD/fQDLmkWWPBhhAAD/94LD/4KC/4KC/76C//+C
w/+Cgv+mgv/Pgv/7gv//gsP/goL/noKCAACCHACCPACCUQCCZQCCeQBBggAAggAAgjwAgoIAQYIA
AIIoAIJNAIJ5AIKCAEEAAAAQEBAgICAwMDBFRUVVVVVlZWV1dXWGhoaampqqqqq6urrLy8vf39/v
7+////9NAABZAABxAACGAACeAAC2AADPAADnAAD/AAD/HBz/NDT/UVH/bW3/ior/oqL/vr5NJABV
KABtNACGPACeSQC2WQDPZQDncQD/fQD/jhz/mjT/plH/sm3/vob/z6L/375NSQBZUQBxaQCGggCe
lgC2rgDPxwDn4wD//wD//xz/+zT/+1H/923/+4b/+6L/+74ATQAAYQAAeQAAjgAApgAAugAA0wAA
6wAA/wAc/xw4/zRV/1Fx/22K/4am/6LD/74AQUEAWVkAcXEAhoYAnp4AtrYAz88A5+cA//9Z//t1
//uK//+e//u6///L///b//8AIEEALFkAOHEARYYAUZ4AXbYAac8AdecAgv8cjv80nv9Rqv9tuv+K
y/+i1/++4/8AAE0AAGUABHkABI4ABKYAAL4AANMAAOsAAP8cJP80PP9RXf9tef+Kkv+iqv++x/8k
AE0wAGVBAIJNAJpZALJlAMtxAOd5AP+CAP+OHP+WNP+mUf+ubf++hv/Lov/bvv9JAE1fAGN1AHqL
AJChAKe3AL3NANTjAOvmF+3qL/DtR/LxX/X0dvf4jvr7pvz/vv8gAAAsAAA4BARJDAhVFBBhIBhx
KCR9OCyGRTiaWU2qbV26gnXLmorfsqLvz77/698gIAA8PABRTQBlWQh5ZQyObRSieRy2fSi+gjjH
jk3PlmHbpnXjso7rw6b308P/69//HBz/HBz/HBz/HBz/HBz/HBz/HBysfHz/HBz/HBz/HBz/HBwA
AABXYnq0tLRtbW1fGku6AAAAD3RFWHRTb2Z0d2FyZQBHcmFmeDKgolNqAAAALklEQVQYlWNgIBIw
AgE6H1kEzEESgTIRIjAGI4JBrAya1STzGdD5dBLAAMRrAQBBhwBFq/1ziwAAAABJRU5ErkJggg==
`)


//scripts
Scripts = map[string][]byte{}

//Load script file to string
Luascript, _ := getContent("./hello.lua") // relative path to make GH version work
Testscript, _ := getContent("./test.lua")
//Luascript := `print("hello WASM from lua")`

//luckily, string to byteslice is easy
//storing as byteslice is future-proof (when the script is actually loaded from file and not from string)
Scripts["hello"] = []byte(Luascript)
Scripts["test"] = []byte(Testscript)


}