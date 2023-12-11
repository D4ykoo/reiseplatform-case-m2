import Axios from "axios-observable";
import { enviroment } from "@/assets/config";

export class UserManagementService {
    
  public getSingleUserRequest(id: number) {
    return Axios.get(enviroment.apiUrl + `/users/${id}`).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }

  public getAllUserRequests() {
    return Axios.get(enviroment.apiUrl + `/users`).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }

  public updateUser(id: number, payload: any) {
    return Axios.put(enviroment.apiUrl + `/users/${id}`, payload).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }

  public deleteUser(id: number) {
    return Axios.delete(enviroment.apiUrl + `/users/${id}`).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }

  public createUser(payload: any) {
    return Axios.post(enviroment.apiUrl + `/users`, payload).subscribe({
      next: (result: any) => {
        console.log(result);
        return result;
      },
      error: (e: any) => {
        console.log(e);
      },
    });
  }
}
