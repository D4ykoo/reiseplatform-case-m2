import { Injectable } from '@angular/core';
import { User } from '../models/user';
import { BehaviorSubject } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment.development';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  private userSubject = new BehaviorSubject<User>({ userid: 0, name: '' });
  public user = this.userSubject.asObservable();

  private currentUser: User | undefined;

  constructor(private httpClient: HttpClient) {}

  checkLoginStatus() {
    
    if (this.userSubject.value.userid == 0) {
      this.httpClient
        .get(environment.Hotel_API + 'loginstatus',{withCredentials:true})
        .subscribe((res) => {
          if (res) {
            this.currentUser = res as User;
            this.userSubject.next(res as User);
          }
        });
    }
  }

  getLoginStatus(): User | undefined {
    return this.currentUser;
  }
}
