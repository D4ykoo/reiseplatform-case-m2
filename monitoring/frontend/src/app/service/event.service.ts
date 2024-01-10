import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class EventService {
  constructor(private httpClient: HttpClient) {}
  public getUserEvents() {
    /* from:Date, to:Date
    let params = new HttpParams();
    params = params.append('from', from.toISOString());
    params = params.append('to', to.toISOString());*/

    return this.httpClient.get(environment.Monitor_API + 'user-events', {
      observe: 'response',
    });
  }

  public getTravelEvents() {
    /* from:Date, to:Date
     let params = new HttpParams();
     params = params.append('from', from.toISOString());
     params = params.append('to', to.toISOString());*/

    return this.httpClient.get(environment.Monitor_API + 'hotel-events', {
      observe: 'response',
    });
  }

  public getCheckoutEvents() {
    /* from:Date, to:Date
     let params = new HttpParams();
     params = params.append('from', from.toISOString());
     params = params.append('to', to.toISOString());*/

    return this.httpClient.get(environment.Monitor_API + 'checkout-events', {
      observe: 'response',
    });
  }
}
