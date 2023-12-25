import { Injectable } from '@angular/core';
import { Hotel } from '../models/hotel';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment.development';
import { BehaviorSubject } from 'rxjs';
import { OfferCart } from '../models/offerCart';

@Injectable({
  providedIn: 'root'
})
export class OfferService {

  private offerSubject = new BehaviorSubject<Map<number, Hotel>>(new Map<number, Hotel>)
  public offers = this.offerSubject.asObservable();

  private offerCart: Map<number, OfferCart> = new Map<number, OfferCart>();

  constructor(private readonly httpClient: HttpClient) { }

  public fetchOffers(): void {
    this.offerSubject.value.clear();
    this.httpClient.get(environment.HotelAPI + "hotels").subscribe((hotelsResponse) => {
      (hotelsResponse as Array<Hotel>).forEach((res: Hotel) => {
        this.offerSubject.value.set(res.id, res)
      })
      this.offerSubject.next(this.offerSubject.getValue());
    })
  }

  public getOffersByHotelId(id: number): Hotel | undefined {
    return this.offerSubject.value.get(id);
  }

  public getCartItem(id: number): OfferCart | undefined {
    return this.offerCart.get(id);
  }

  public getCart(): Array<OfferCart | undefined> {
    return Array.from(this.offerCart.values());
  }


  public addToCart(offer: OfferCart) {
    this.offerCart.set(offer.travelId,offer);
  }

  private handleError(error: any): Promise<any> {
    console.log('Interner Fehler: ' + error.message);
    return Promise.reject(error.message || error);
  }
}


