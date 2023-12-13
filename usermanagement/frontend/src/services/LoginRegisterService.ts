import Axios from "axios-observable";
import { enviroment } from "@/assets/config";
import type { LoginUser, RegisterUser, ResetUser, User } from "@/models/UserModel";
import { throwError, type Observable, catchError } from "rxjs";
import type { AxiosError, AxiosResponse } from "axios";
import { UtilService } from "./UtilService";

export class LoginRegisterService {
  utils = new UtilService();

  public LoginRequest(payload: LoginUser): Observable<AxiosResponse<any>> {
    console.log(enviroment.apiUrl + `/login`);
    console.log(this.utils.getJwtFromToken());
    return Axios.post(enviroment.apiUrl + "/login", payload).pipe(catchError(this.handleError))
  }

  public RegisterRequest(payload: RegisterUser): Observable<AxiosResponse<any>> {
    console.log(enviroment.apiUrl + `/login`);
    return Axios.post(enviroment.apiUrl + `/register`, payload).pipe(catchError(this.handleError))
  }
  
  public ResetPasswordRequest(payload: ResetUser): Observable<AxiosResponse<any>> {
    return Axios.put(enviroment.apiUrl + `/reset`, payload).pipe(catchError(this.handleError))
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
