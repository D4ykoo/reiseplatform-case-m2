import Axios from "axios-observable";
import { enviroment, headerConf } from "@/assets/config";
import type { LoginUser, RegisterUser, ResetUser } from "@/models/UserModel";
import { throwError, type Observable, catchError } from "rxjs";
import type { AxiosError, AxiosResponse } from "axios";
import { UtilService } from "./UtilService";

export class LoginRegisterService {
  utils = new UtilService();

  public LoginRequest(payload: LoginUser): Observable<AxiosResponse<any>> {
    return Axios.post(enviroment.apiUrl + "/login", payload, headerConf).pipe(catchError(this.handleError))
  }

  public RegisterRequest(payload: RegisterUser): Observable<AxiosResponse<any>> {
    return Axios.post(enviroment.apiUrl + `/register`, payload, headerConf).pipe(catchError(this.handleError))
  }
  
  public ResetPasswordRequest(payload: ResetUser): Observable<AxiosResponse<any>> {
    return Axios.put(enviroment.apiUrl + `/reset`, payload, headerConf).pipe(catchError(this.handleError))
  }

  public LogoutRequest(): Observable<AxiosResponse<any>> {
    return Axios.get(enviroment.apiUrl + `/logout`,headerConf).pipe(catchError(this.handleError))
  }

  private handleError(error: AxiosError){
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.message);
    } else {
        // The backend returned an unsuccessful response code.
        // The response body may contain clues as to what went wrong.
        console.error(
            `Backend returned code ${error.status}, body was: `, error.message);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Unkown error'));

}

}
