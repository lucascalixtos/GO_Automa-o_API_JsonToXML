package main

import(
    "crypto/tls"
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "os"
    "encoding/json"
    "strings"
    "encoding/xml"
    //"strconv"
)

type AccountsBase struct{
    AccountsBase []Accounts `json:"accounts"`
}

type Accounts struct {
    Nome    string   `json:"name"`
    CDL     string   `json:"userVars.CDL"`
    Conta1  string   `json:"userVars.MT940Cta1"`
    Conta2  string   `json:"userVars.MT940Cta2"`
    Conta3  string   `json:"userVars.MT940Cta3"`
    Conta4  string   `json:"userVars.MT940Cta4"`
    Conta5  string   `json:"userVars.MT940Cta5"`
    Conta6  string   `json:"userVars.MT940Cta6"`
    Conta7  string   `json:"userVars.MT940Cta7"`
    Conta8  string   `json:"userVars.MT940Cta8"`
    Conta9  string   `json:"userVars.MT940Cta9"`
    Conta10 string   `json:"userVars.MT940Cta10"`
    Conta11 string   `json:"userVars.MT940Cta11"`
    Conta12 string   `json:"userVars.MT940Cta12"`
    Conta13 string   `json:"userVars.MT940Cta13"`
    Conta14 string   `json:"userVars.MT940Cta14"`
    Conta15 string   `json:"userVars.MT940Cta15"`
    Conta16 string   `json:"userVars.MT940Cta16"`
    Conta17 string   `json:"userVars.MT940Cta17"`
    Conta18 string   `json:"userVars.MT940Cta18"`
    Conta19 string   `json:"userVars.MT940Cta19"`
    Conta20 string   `json:"userVars.MT940Cta20"`
}

type XmlAccounts struct{
    XMLName     xml.Name    `xml:"MT940RPT,omitempty"`
    CDL         string      `xml:"Header>CDL,omitempty"`
    Value       [20]getId   `xml:"MT940ACCT,omitempty"` 
}

type getId struct{
    ID string `xml:"ID,omitempty"`
} 

type Conf struct{
    Username    string  `json:"user"` 
    Passwd      string  `json:"password"` 
    Link        string  `json:"link"`
    DirXML      string  `json:"dirXML"`
    DirLOG      string  `json:"dirLOG"`
    DirAccounts string  `json:"dirAccounts"`
}

func basicAuth() string {
    
    jsonFile, err := os.Open(`MT940gen.conf`)

    if err != nil {
        //Caso tenha tido erro, ele é apresentado na tela
        fmt.Println(err)
    }

    defer jsonFile.Close()

    //Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
    byteValueJSON, _:= ioutil.ReadAll(jsonFile)

    //Declaração abreviada de um objeto do tipo Conf
    conf := Conf{}

    //Conversão da variável byte em um objeto do tipo struct Conf
    json.Unmarshal(byteValueJSON, &conf)

    var username string = conf.Username
    var passwd string = conf.Passwd

    transCfg := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // disable verify
    }
    client := &http.Client{Transport: transCfg}
    req, err := http.NewRequest("GET", conf.Link, nil)
    req.SetBasicAuth(username, passwd)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)
    s := string(bodyText)

    f, err := os.Create(conf.DirAccounts + "accounts.json")
    if err != nil {
        fmt.Println(err)
    }
    
    l, err := f.WriteString(s)
    if err != nil {
        fmt.Println(err)
        f.Close()
    }

    fmt.Println(l, "Arquivo criado")
    err = f.Close()
    return s
}

func main(){
    jsonFile, err := os.Open(`MT940gen.conf`)

    if err != nil {
        //Caso tenha tido erro, ele é apresentado na tela
        fmt.Println(err)
    }

    defer jsonFile.Close()

    //Aqui o arquivo é convertido para uma variável array de bytes, através do pacote "io/ioutil"
    byteValueJSON, _:= ioutil.ReadAll(jsonFile)

    //Declaração abreviada de um objeto do tipo Conf
    conf := Conf{}

    //Conversão da variável byte em um objeto do tipo struct Conf
    json.Unmarshal(byteValueJSON, &conf)

    dirLog := conf.DirLOG

    f, err := os.OpenFile(dirLog + "MT940gen.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer f.Close()

    log.SetOutput(f)

    fmt.Println("Requesting...")
    S := basicAuth()
    fmt.Println(S)

    data := AccountsBase{}

    file, _ := ioutil.ReadFile(conf.DirAccounts + "accounts.json")

    _ = json.Unmarshal([]byte(file), &data)
    save := &XmlAccounts{CDL: "teste"}

    for i := 0; i < len(data.AccountsBase); i++ {
        if !strings.HasPrefix("userVars.CDL", data.AccountsBase[i].CDL) {
            fmt.Println(data.AccountsBase[i].Nome)
            fmt.Println("CDL: ", data.AccountsBase[i].CDL)
            
            log.Print("Account Name: " + data.AccountsBase[i].Nome)
            log.Print("CDL: " + data.AccountsBase[i].CDL)
            
            save = &XmlAccounts{CDL: data.AccountsBase[i].CDL}
            j := 0

            if !strings.HasPrefix("userVars.MT940Cta1", data.AccountsBase[i].Conta1){
                fmt.Println("ID: ", data.AccountsBase[i].Conta1)
                save.Value[j].ID = data.AccountsBase[i].Conta1
                log.Print("ID: " + data.AccountsBase[i].Conta1)
                j++
            }

            if !strings.HasPrefix("userVars.MT940Cta2", data.AccountsBase[i].Conta2){
                fmt.Println("ID: ", data.AccountsBase[i].Conta2)
                save.Value[j].ID = data.AccountsBase[i].Conta2
                log.Print("ID: " + data.AccountsBase[i].Conta2)
                j++
            }

            if !strings.HasPrefix("userVars.MT940Cta3", data.AccountsBase[i].Conta3){
                fmt.Println("ID: ", data.AccountsBase[i].Conta3)
                save.Value[j].ID = data.AccountsBase[i].Conta3
                log.Print("ID: " + data.AccountsBase[i].Conta3)
                j++
            }

            if !strings.HasPrefix("userVars.MT940Cta4", data.AccountsBase[i].Conta4){
                fmt.Println("ID: ", data.AccountsBase[i].Conta4)
                save.Value[j].ID = data.AccountsBase[i].Conta4
                log.Print("ID: " + data.AccountsBase[i].Conta4)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta5", data.AccountsBase[i].Conta5){
                fmt.Println("ID: ", data.AccountsBase[i].Conta5)
                save.Value[j].ID = data.AccountsBase[i].Conta5
                log.Print("ID: " + data.AccountsBase[i].Conta5)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta6", data.AccountsBase[i].Conta6){
                fmt.Println("ID: ", data.AccountsBase[i].Conta6)
                save.Value[j].ID = data.AccountsBase[i].Conta6
                log.Print("ID: " + data.AccountsBase[i].Conta6)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta7", data.AccountsBase[i].Conta7){
                fmt.Println("ID: ", data.AccountsBase[i].Conta7)
                save.Value[j].ID = data.AccountsBase[i].Conta7
                log.Print("ID: " + data.AccountsBase[i].Conta7)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta8", data.AccountsBase[i].Conta8){
                fmt.Println("ID: ", data.AccountsBase[i].Conta8)
                save.Value[j].ID = data.AccountsBase[i].Conta8
                log.Print("ID: " + data.AccountsBase[i].Conta8)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta9", data.AccountsBase[i].Conta9){
                fmt.Println("ID: ", data.AccountsBase[i].Conta9)
                save.Value[j].ID = data.AccountsBase[i].Conta9
                log.Print("ID: " + data.AccountsBase[i].Conta9)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta10", data.AccountsBase[i].Conta10){
                fmt.Println("ID: ", data.AccountsBase[i].Conta10)
                save.Value[j].ID = data.AccountsBase[i].Conta10
                log.Print("ID: " + data.AccountsBase[i].Conta10)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta11", data.AccountsBase[i].Conta11){
                fmt.Println("ID: ", data.AccountsBase[i].Conta11)
                save.Value[j].ID = data.AccountsBase[i].Conta11
                log.Print("ID: " + data.AccountsBase[i].Conta11)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta12", data.AccountsBase[i].Conta12){
                fmt.Println("ID: ", data.AccountsBase[i].Conta12)
                save.Value[j].ID = data.AccountsBase[i].Conta12
                log.Print("ID: " + data.AccountsBase[i].Conta12)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta13", data.AccountsBase[i].Conta13){
                fmt.Println("ID: ", data.AccountsBase[i].Conta13)
                save.Value[j].ID = data.AccountsBase[i].Conta13
                log.Print("ID: " + data.AccountsBase[i].Conta13)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta14", data.AccountsBase[i].Conta14){
                fmt.Println("ID: ", data.AccountsBase[i].Conta14)
                save.Value[j].ID = data.AccountsBase[i].Conta14
                log.Print("ID: " + data.AccountsBase[i].Conta14)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta15", data.AccountsBase[i].Conta15){
                fmt.Println("ID: ", data.AccountsBase[i].Conta15)
                save.Value[j].ID = data.AccountsBase[i].Conta15
                log.Print("ID: " + data.AccountsBase[i].Conta15)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta16", data.AccountsBase[i].Conta16){
                fmt.Println("ID: ", data.AccountsBase[i].Conta16)
                save.Value[j].ID = data.AccountsBase[i].Conta16
                log.Print("ID: " + data.AccountsBase[i].Conta16)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta17", data.AccountsBase[i].Conta17){
                fmt.Println("ID: ", data.AccountsBase[i].Conta17)
                save.Value[j].ID = data.AccountsBase[i].Conta17
                log.Print("ID: " + data.AccountsBase[i].Conta17)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta18", data.AccountsBase[i].Conta18){
                fmt.Println("ID: ", data.AccountsBase[i].Conta18)
                save.Value[j].ID = data.AccountsBase[i].Conta18
                log.Print("ID: " + data.AccountsBase[i].Conta18)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta19", data.AccountsBase[i].Conta19){
                fmt.Println("ID: ", data.AccountsBase[i].Conta19)
                save.Value[j].ID = data.AccountsBase[i].Conta19
                log.Print("ID: " + data.AccountsBase[i].Conta19)
                j++  
            }

            if !strings.HasPrefix("userVars.MT940Cta20", data.AccountsBase[i].Conta20){
                fmt.Println("ID: ", data.AccountsBase[i].Conta20)
                save.Value[j].ID = data.AccountsBase[i].Conta20
                log.Print("ID: " + data.AccountsBase[i].Conta20)
                j++  
            }

            log.Print("-------------")

            file, _ := xml.MarshalIndent(save, "", " ")

            path := conf.DirXML
            
            if _, err := os.Stat(path); os.IsNotExist(err) {
                os.MkdirAll(path, os.ModePerm)
            }
            fileName :=  conf.DirXML + "MT940-" + strings.TrimSpace(data.AccountsBase[i].Nome) + ".xml"
            _ = ioutil.WriteFile(fileName, file, 0644)
        }
	}

    log.Print("Directory Accounts file: " + conf.DirAccounts)
    log.Print("Directory XML files: " + conf.DirXML)

    log.Print("------------------------------------------------------------")


    //fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))


}