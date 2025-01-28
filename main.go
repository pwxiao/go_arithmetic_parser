package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
    "unicode"
)

// Token 类型
type TokenType int

const (
    INVALID TokenType = iota
    NUMBER
    PLUS
    MINUS
    MULTIPLY
    DIVIDE
    LPAREN
    RPAREN
    POWER
    SQRT
)

// Token 结构体
type Token struct {
    Type  TokenType
    Value string
}

// Lexer 结构体
type Lexer struct {
    Text        string
    Position    int
    CurrentChar byte
}

// 创建新的 Lexer
func NewLexer(text string) *Lexer {
    return &Lexer{Text: text, Position: 0, CurrentChar: text[0]}
}

// 前进到下一个字符
func (l *Lexer) Advance() {
    l.Position++
    if l.Position >= len(l.Text) {
        l.CurrentChar = 0
    } else {
        l.CurrentChar = l.Text[l.Position]
    }
}

// 跳过空白字符
func (l *Lexer) SkipWhitespace() {
    for l.CurrentChar != 0 && unicode.IsSpace(rune(l.CurrentChar)) {
        l.Advance()
    }
}

// 获取数字（支持浮点数）
func (l *Lexer) GetNumber() string {
    var result strings.Builder
    hasDot := false
    for l.CurrentChar != 0 && (unicode.IsDigit(rune(l.CurrentChar)) || l.CurrentChar == '.') {
        if l.CurrentChar == '.' {
            if hasDot {
                break
            }
            hasDot = true
        }
        result.WriteByte(l.CurrentChar)
        l.Advance()
    }
    return result.String()
}

// 获取下一个 Token
func (l *Lexer) GetNextToken() Token {
    for l.CurrentChar != 0 {
        if unicode.IsSpace(rune(l.CurrentChar)) {
            l.SkipWhitespace()
            continue
        }

        if unicode.IsDigit(rune(l.CurrentChar)) || l.CurrentChar == '.' {
            return Token{Type: NUMBER, Value: l.GetNumber()}
        }

        switch l.CurrentChar {
        case '+':
            l.Advance()
            return Token{Type: PLUS, Value: "+"}
        case '-':
            l.Advance()
            return Token{Type: MINUS, Value: "-"}
        case '*':
            l.Advance()
            return Token{Type: MULTIPLY, Value: "*"}
        case '/':
            l.Advance()
            return Token{Type: DIVIDE, Value: "/"}
        case '(':
            l.Advance()
            return Token{Type: LPAREN, Value: "("}
        case ')':
            l.Advance()
            return Token{Type: RPAREN, Value: ")"}
        case '^':
            l.Advance()
            return Token{Type: POWER, Value: "^"}
        case 's':
            if strings.HasPrefix(l.Text[l.Position:], "sqrt(") {
                l.Position += 4
                l.CurrentChar = l.Text[l.Position]
                return Token{Type: SQRT, Value: "sqrt"}
            }
        default:
            return Token{Type: INVALID, Value: string(l.CurrentChar)}
        }
    }
    return Token{Type: INVALID, Value: ""}
}

// Parser 结构体
type Parser struct {
    Lexer        *Lexer
    CurrentToken Token
}

// 创建新的 Parser
func NewParser(lexer *Lexer) *Parser {
    token := lexer.GetNextToken()
    return &Parser{Lexer: lexer, CurrentToken: token}
}

// 消费当前 Token 并获取下一个 Token
func (p *Parser) Eat(tokenType TokenType) {
    if p.CurrentToken.Type == tokenType {
        p.CurrentToken = p.Lexer.GetNextToken()
    } else {
        panic(fmt.Sprintf("Unexpected token: %v", p.CurrentToken))
    }
}

// 解析因子
func (p *Parser) ParseFactor() float64 {
    token := p.CurrentToken
    if token.Type == NUMBER {
        p.Eat(NUMBER)
        value, _ := strconv.ParseFloat(token.Value, 64)
        return value
    } else if token.Type == LPAREN {
        p.Eat(LPAREN)
        result := p.ParseExpression()
        p.Eat(RPAREN)
        return result
    } else if token.Type == SQRT {
        p.Eat(SQRT)
        p.Eat(LPAREN)
        result := math.Sqrt(p.ParseExpression())
        p.Eat(RPAREN)
        return result
    } else {
        panic(fmt.Sprintf("Unexpected token: %v", token))
    }
}

// 解析项
func (p *Parser) ParseTerm() float64 {
    result := p.ParseFactor()
    for p.CurrentToken.Type == MULTIPLY || p.CurrentToken.Type == DIVIDE {
        token := p.CurrentToken
        if token.Type == MULTIPLY {
            p.Eat(MULTIPLY)
            result *= p.ParseFactor()
        } else if token.Type == DIVIDE {
            p.Eat(DIVIDE)
            result /= p.ParseFactor()
        }
    }
    return result
}

// 解析表达式
func (p *Parser) ParseExpression() float64 {
    result := p.ParseTerm()
    for p.CurrentToken.Type == PLUS || p.CurrentToken.Type == MINUS || p.CurrentToken.Type == POWER {
        token := p.CurrentToken
        if token.Type == PLUS {
            p.Eat(PLUS)
            result += p.ParseTerm()
        } else if token.Type == MINUS {
            p.Eat(MINUS)
            result -= p.ParseTerm()
        } else if token.Type == POWER {
            p.Eat(POWER)
            result = math.Pow(result, p.ParseTerm())
        }
    }
    return result
}

// 主函数
func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("输入表达式: ")
        expression, _ := reader.ReadString('\n')
        expression = strings.TrimSpace(expression)
        if expression == "exit" {
            break
        }
        lexer := NewLexer(expression)
        parser := NewParser(lexer)
        result := parser.ParseExpression()
        fmt.Printf("结果: %f\n", result)
    }
}