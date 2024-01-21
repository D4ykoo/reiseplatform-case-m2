import { Injectable } from '@angular/core';
import { environment } from '../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class CartService {
  apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) {}

  public getCart() {
    console.log(this.apiUrl);
    return this.http.get(`${this.apiUrl}/cart/find`);
  }

  public removeFromCart(cartId: number, hotelId: number, travelId: number) {
    // ttp://localhost:8084/api/v1/cart/entry/{cart_id}/{hotel_id}/{travel_id}
    return this.http.delete(
      `${this.apiUrl}/cart/entry/${cartId}/${hotelId}/${travelId}`,
    );
  }
}
