package main

import (
//	"fmt"
	"log"
	)

type node struct {
	kind       string
	value      string
	name       string
//	callee     *node
//	expression *node
	body       []node
//	params     []node
//	arguments  *[]node
//	context    *[]node
}

type tree node
var x node
var pc int
var pt []token

func parser(tokens []token) tree {
	pc = 0
	pt = tokens

	tree := tree{
		kind: "TREE",
		body: []node{},
	}
	for pc < len(pt) {
		tree.body = append(tree.body, walk())
	}
	return tree
}

func walk() node {
	n := node{}

	//PROGRAM
	if pt[pc].kind == "reserved" && pt[pc].value == "program" {
		pc++
		if pt[pc].kind == "llave" && pt[pc].value == "{" {
			pc++
			x = walk()
			if x.kind == "lista-declaracion"{
				n = node{
					kind: "program",
					body: []node{},
				}
				n.body = append(n.body, x)
				/*pc++
				x := walk()
				if x.kind == "lista-sentencias" {
					n.body = append(n.body, x)
					pc++*/
					if pt[pc].kind == "llave" && pt[pc].value == "}" {
						pc++
						return n
					}
				//}
			}
		}
	}

	//LISTA-DECLARACION
	if x.kind == "declaracion" {
		x = n
		n = node{
			kind: "lista-declaracion",
			body: []node{},
		}
		for x.kind == "declaracion" {
			n.body = append(n.body, x)
			pc++
			x = walk()
		}
		pc++
		return n
	}

	//DECLARACION
	if x.kind == "tipo" {
		x = walk()
		if x.kind == "lista-id" {
			n = node{
				kind: "declaracion",
				body: []node{},
			}
			n.body = append(n.body, x)
			if pt[pc].kind == "p-comma" {
				x = node{
					kind: "p-comma",
				}
				n.body = append(n.body, x)
				return n
			}
		}
	}

	//TIPO
	if pt[pc].kind == "reserved" && (pt[pc].value == "int" || pt[pc].value == "float" || pt[pc].value == "bool") {
		return node {
			kind: "tipo",
			value: pt[pc].value,
		}
	}

	//LISTA-ID
	if pt[pc].kind == "identificador" {
		pc++
		for pt[pc].kind == "comma" && pt[pc+1].kind == "identificador"  && pc+2 < len(pt) {
			pc+=2
		}
		return node {
			kind: "lista-id",
			value: "" + pt[pc].kind,
		}
	}

	//LISTA-SENTENCIAS
	if x.kind == "sentencia" {
		x = n
		n = node{
			kind: "lista-sentencias",
			body: []node{},
		}
		for x.kind == "sentencia" {
			n.body = append(n.body, x)
			pc++
			x = walk()
		}
		pc++
		return n
	}

	//SENTENCIA
	if x.kind == "seleccion" || x.kind == "iteracion" || x.kind == "repeticion" || x.kind == "sent-read" || x.kind == "sent-write" || x.kind == "bloque" || x.kind == "asignacion" {
		return node {
			kind: "sentencia",
			value: x.kind,
		}
	}

	//SELECCION
	if pt[pc].kind == "reserved" && pt[pc].value == "if" {
		x = node{ kind: "if", }
		n = node{ kind: "seleccion", body: []node{}, }
		n.body = append(n.body, x)
		if pt[pc+1].kind == "paren" && pt[pc+1].value == "("{
			x = node{ kind: "paren", value: "(", }
			n.body = append(n.body, x)
			x = walk()
			if x.kind == "expresion" {
				n.body = append(n.body, x)
				if pt[pc].kind == "paren" && pt[pc].value == ")"{
					x = node{ kind: "paren", value: ")", }
					n.body = append(n.body, x)
					x = walk()
					if x.kind == "bloque" {
						n.body = append(n.body, x)
						x = walk()
						if x.kind == "reserved" && x.value == "else"{
							n.body = append(n.body, x)
							x = walk()
							if x.kind == "bloque" {
								n.body = append(n.body, x)
								x = walk()
							}
						}
						if x.kind == "reserved" && x.value == "fi"{
							n.body = append(n.body, x)
							return n
						}
					}
				}
			}
		}
	}

	//ITERACION
	if pt[pc].kind == "reserved" && pt[pc].value == "while" {
		x = node{ kind: "while", }
		n = node{ kind: "iteracion", body: []node{}, }
		n.body = append(n.body, x)
		if pt[pc+1].kind == "paren" && pt[pc+1].value == "("{
			x = node{ kind: "paren", value: "(", }
			n.body = append(n.body, x)
			x = walk()
			if x.kind == "expresion" {
				n.body = append(n.body, x)
				if pt[pc].kind == "paren" && pt[pc].value == ")"{
					x = node{ kind: "paren", value: ")", }
					n.body = append(n.body, x)
					x = walk()
					if x.kind == "bloque" {
						n.body = append(n.body, x)
						return n
					}
				}
			}
		}
	}

	//REPETICION
	if pt[pc].kind == "reserved" && pt[pc].value == "do" {
		x = node{ kind: "do", }
		n = node{ kind: "repeticion", body: []node{}, }
		n.body = append(n.body, x)
		if x.kind == "bloque" {
			n.body = append(n.body, x)
			x = walk()
			pc++
			if pt[pc].kind == "reserved" && pt[pc].value == "until" {
				x = node{ kind: "do", }
				n.body = append(n.body, x)
				pc++
				if pt[pc].kind == "paren" && pt[pc].value == "(" {
					x = node{ kind: "paren", value: "("}
					n.body = append(n.body, x)
					x = walk()
					if x.kind == "expresion" {
						n.body = append(n.body, x)
						pc++
						if pt[pc].kind == "paren" && pt[pc].value == ")" {
							x = node{ kind: "paren", value: ")"}
							n.body = append(n.body, x)
							pc++
							if pt[pc].kind == "p-comma" {
								x = node{ kind: "p-comma", value: ";"}
								n.body = append(n.body, x)
								pc++
								return n
							}
						}
					}
				}
			}
		}
	}

	//SENT-READ
	if pt[pc].kind == "reserved" && pt[pc].value == "read" {
		x = node{ kind: "read", }
		n = node{ kind: "sent-read", body: []node{}, }
		n.body = append(n.body, x)
		pc++
		if pt[pc].kind == "identificador" {
			x = node{ kind: "identificador", value: pt[pc].value, }
			n.body = append(n.body, x)
			pc++
			if pt[pc].kind == "p-comma" {
				x = node{ kind: "p-comma", value: ";"}
				n.body = append(n.body, x)
				pc++
				return n
			}
		}
	}

	//SENT-WRITE
	if pt[pc].kind == "reserved" && pt[pc].value == "write" {
		x = node{ kind: "write", }
		n = node{ kind: "sent-write", body: []node{}, }
		n.body = append(n.body, x)
		x = walk()
		if x.kind == "expresion" {
			n.body = append(n.body, x)
			pc++
			if pt[pc].kind == "p-comma" {
				x = node{ kind: "p-comma", value: ";"}
				n.body = append(n.body, x)
				pc++
				return n
			}
		}
	}

	//BLOQUE
	if pt[pc].kind == "llave" && pt[pc].value == "{" {
		x = node{ kind: "llave", value: "{", }
		n = node{ kind: "bloque", body: []node{}, }
		n.body = append(n.body, x)
		x = walk()
		if x.kind == "lista-sentencias" {
			n.body = append(n.body, x)
			pc++
			if pt[pc].kind == "llave" && pt[pc].value == "}" {
				x = node{ kind: "llave", value: "}"}
				n.body = append(n.body, x)
				pc++
				return n
			}
		}
	}

	//ASIGNACION
	if pt[pc].kind == "identificador" {
		n = node { kind: "asignacion", }
		pc++
		x = walk()
		if pt[pc].kind == "equals" && x.kind == "expresion" {
			n.body = append(n.body, node{ kind: "equals", value: "="})
			n.body = append(n.body, x)
			if pt[pc].kind == "p-comma" {
				n.body = append(n.body, node{ kind: "p-comma", value: ";"})
				pc++
				return n
			}
		}
	}

	//EXPRESION
	if x.kind == "term-and" {
		pc++
		x = walk()
		for pt[pc].kind == "reserved" && pt[pc].value == "or"  && x.kind == "term-and" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "expresion",
		}
	}

	//TERM-AND
	if x.kind == "igualdad" {
		pc++
		x = walk()
		for pt[pc].kind == "reserved" && pt[pc].value == "and"  && x.kind == "igualdad" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "term-and",
		}
	}

	//IGUALDAD
	if x.kind == "relacion" {
		pc++
		x = walk()
		for (pt[pc].kind == "comparator" && (pt[pc].value == "==" || pt[pc].value == "!="))  && x.kind == "relacion" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "igualdad",
		}
	}

	//RELACION
	if x.kind == "expresion-suma" {
		pc++
		x = walk()
		for (pt[pc].kind == "comparator" && (pt[pc].value == "<" || pt[pc].value == "<=" || pt[pc].value == ">" || pt[pc].value == ">="))  && x.kind == "expresion-suma" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "relacion",
		}
	}

	//EXPRESION-SUMA
	if x.kind == "termino" {
		pc++
		x = walk()
		for (pt[pc].kind == "op-suma" && (pt[pc].value == "-" || pt[pc].value == "+"))  && x.kind == "termino" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "expresion-suma",
		}
	}

	//TERMINO
	if x.kind == "unario" {
		pc++
		x = walk()
		for (pt[pc].kind == "op-mult" && (pt[pc].value == "/" || pt[pc].value == "*"))  && x.kind == "unario" && pc+2 < len(pt) {
			pc++
			x = walk()
		}
		return node {
			kind: "termino",
		}
	}

	//UNARIO
	if (pt[pc].kind == "op-suma" && (pt[pc].value == "-" || pt[pc].value == "+")) || (pt[pc].kind == "reserved" && pt[pc].value == "not") { 
		x = walk()
		if x.kind == "factor" {
			return node {
				kind: "unario",
			}
		}
	}

	//FACTOR
	if (pt[pc].kind == "paren" && pt[pc].value == "(") || pt[pc].kind == "identificador"  || pt[pc].kind == "numero" || (pt[pc].kind == "reserved" && (pt[pc].value == "true" || pt[pc].value == "true")){
		x = node{ kind: "paren", value: "(", }
		n = node{ kind: "factor", body: []node{}, }
		n.body = append(n.body, x)
		x = walk()
		if x.kind == "expresion" {
			n.body = append(n.body, x)
			pc++
			if pt[pc].kind == "paren" && pt[pc].value == ")" {
				x = node{ kind: "paren", value: ")"}
				n.body = append(n.body, x)
				pc++
			}
		}
		return node{
			kind: "factor",
		}
	}

	//os.Exit(1)
	log.Fatal(pt[pc].kind)
	return node{}
}
