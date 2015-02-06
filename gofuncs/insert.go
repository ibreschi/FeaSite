package hallo

import (
	"net/http"
)


func processRequest (r *http.Request,title string,param *Params) error {
    
    switch title {
        case "insertedValue":
            return processInsertedValueData(r,param)
        case "login":
            return processLoginFunc(r,param)
        default:
            return nil
        }
}


func insert(w http.ResponseWriter, r *http.Request, title string) {

    if r.Method == "GET" {
        http.Error(w, "NOT POSSIBLE", http.StatusInternalServerError)
    } else {
               
        myfunc, err := insertFuncManager.getFunction(title)
        if err!=nil {
            // no insert function for this title
            // curl --data "param1=2" http://localhost:8080/insert/MAO
            renderTemplate(w,"Error",err)
            return
        }

        param, err := myfunc(r)
        if err!=nil {
            // no params fetched with the insert function 
            // curl --data "username=&&password=2&&token=3" http://localhost:8080/insert/login
            renderTemplate(w,"Error",err)
            return
        }

        // get token to avoid double summit 
        r.ParseForm() 

        if token := r.Form.Get("token"); token == ""{
            renderTemplate(w,"Error","duplicate submission")
            return
        }


            //request, err := insertRequestManager.getRequest(title)
            //err := request(r, param)

        err = processRequest(r,title,param)
        if err!=nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }else {
            renderTemplate(w,title,param)
        }
    }
}