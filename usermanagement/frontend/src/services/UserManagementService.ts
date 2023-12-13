import Axios from "axios-observable";
import { enviroment } from "@/assets/config";
import type { UpdateUser, User } from "@/models/UserModel";
import { throwError, type Observable, catchError } from "rxjs";
import type { AxiosError, AxiosResponse } from "axios";
import { UtilService } from "./UtilService";

export class UserManagementService {

  utils = new UtilService(); 

  headerConf = {
    headers: {
      Authorization: `Bearer `,
    },
    withCredentials: true 
  }

  public getSingleUserRequest(id: number) {
    this.headerConf.headers.Authorization = `Bearer ${this.utils.getJwtFromToken()}`

    return Axios.get(enviroment.apiUrl + `/users/${id}`, this.headerConf).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }

  public getAllUserRequests(): Observable<AxiosResponse<any>> {
    this.headerConf.headers.Authorization = `Bearer ${this.utils.getJwtFromToken()}`
    console.log(this.headerConf);
    return Axios.get(enviroment.apiUrl + `/users`, this.headerConf).pipe(
      catchError(this.handleError)
    );
  }

  public updateUser(id: number, payload: UpdateUser) {
    return Axios.put(enviroment.apiUrl + `/users/${id}`, payload).pipe(
      catchError(this.handleError)
    );
  }

  public deleteUser(id: number) {
    return Axios.delete(enviroment.apiUrl + `/users/${id}`).pipe(
      catchError(this.handleError)
    );
  }

  public createUser(payload: any) {
    return Axios.post(enviroment.apiUrl + `/users`, payload).pipe(
      catchError(this.handleError)
    );
  }

  private handleError(error: AxiosError) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error("An error occurred:", error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `,
        error.message
      );
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error("Unkown error"));
  }
}
