import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class UserService {
  constructor(private httpClient: HttpClient) {}
  public getUserEvents() {
    /* from:Date, to:Date
    let params = new HttpParams();
    params = params.append('from', from.toISOString());
    params = params.append('to', to.toISOString());*/

    return this.httpClient.get('http://localhost:3000/api/v1/user-events', {
      observe: 'response',
    });
  }

  public getTravelEvents() {
    /* from:Date, to:Date
     let params = new HttpParams();
     params = params.append('from', from.toISOString());
     params = params.append('to', to.toISOString());*/

    return this.httpClient.get('http://localhost:3000/api/v1/hotel-events', {
      observe: 'response',
    });
  }

  public getCheckoutEvents() {
    /* from:Date, to:Date
     let params = new HttpParams();
     params = params.append('from', from.toISOString());
     params = params.append('to', to.toISOString());*/

    return this.httpClient.get('http://localhost:3000/api/v1/checkout-events', {
      observe: 'response',
    });
  }
}
