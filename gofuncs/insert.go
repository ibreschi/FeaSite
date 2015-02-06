package hallo

import (
	"net/http"
)


// func getParams (title string,r *http.Request) *Params {
//     switch title {
//         case "insertedValue":
//            // return retriveInsertedValueData(r)
//         case "login":
//            // return retriveLoginData(r)
//         default:
//             return nil
//         }
//}


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
       
        r.ParseForm()     
        
        myfunc, err := insertFuncManager.getFunction(title)
        // 
        if err!=nil {
            // no insert function for this title
            return
        }
        param, err := myfunc(r)

        if err!=nil {
            //param := getParams(title,r)
            token := r.Form.Get("token")
            if  token != "" && param != nil  {
                err := processRequest(r,title,param)
                if err!=nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                }else {
                    renderTemplate(w,title,param)
                }
            } else {
                // error 
                // duplicate submission 
                // empty value
                renderTemplate(w,title+"Error",nil)
            }
        }
    }
}