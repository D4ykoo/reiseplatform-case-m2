import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class CartService {
  apiUrl = environment.api;

  constructor(private http: HttpClient) {}

  public getCart(id: number) {
    console.log(id);
    console.log(`${this.apiUrl}/cart/${id}`);
    return this.http.get(`${this.apiUrl}/cart/${id}`);
  }

  public addToCart(payload: {userId: number, hotelId: number, travelId: number}) {
    return this.http.post(`${this.apiUrl}/cart`, payload).subscribe((res) => {
      console.log(res);
    });
  }

  public removeFromCart(cartId: number, hotelId: number, travelId: number) {
    // ttp://localhost:8084/api/v1/cart/entry/{cart_id}/{hotel_id}/{travel_id}
    return this.http.delete(`${this.apiUrl}/cart/entry/${cartId}/${hotelId}/${travelId}`);
  }

}
