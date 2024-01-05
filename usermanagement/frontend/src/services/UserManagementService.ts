import Axios from "axios-observable";
import { enviroment, headerConf } from "@/assets/config";
import type { UpdateUser } from "@/models/UserModel";
import { throwError, type Observable, catchError } from "rxjs";
import type { AxiosError, AxiosResponse } from "axios";
import { UtilService } from "./UtilService";

export class UserManagementService {
  utils = new UtilService();

  apiUrl = import.meta.env.VITE_API_URL;

  public getSingleUserRequest(id: number) {
    return Axios.get(this.apiUrl + `/users/${id}`, headerConf).subscribe({
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
    return Axios.get(this.apiUrl + `/users`, headerConf).pipe(
      catchError(this.handleError),
    );
  }

  public updateUser(id: number, payload: UpdateUser) {
    return Axios.put(this.apiUrl + `/users/${id}`, payload, headerConf).pipe(
      catchError(this.handleError),
    );
  }

  public deleteUser(id: number) {
    return Axios.delete(this.apiUrl + `/users/${id}`, headerConf).pipe(
      catchError(this.handleError),
    );
  }

  public createUser(payload: any) {
    return Axios.post(this.apiUrl + `/users`, payload, headerConf).pipe(
      catchError(this.handleError),
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
        error.message,
      );
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error("Unkown error"));
  }
}
