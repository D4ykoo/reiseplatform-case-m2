import { Injectable } from '@angular/core';
import { Hotel, Hoteldemo } from '../models/hotel';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../environments/environment.development';
import { BehaviorSubject } from 'rxjs';
import { OfferCart } from '../models/offerCart';
import { Tag } from '../models/tag';


@Injectable({
  providedIn: 'root'
})
export class OfferService {

  private offerSubject = new BehaviorSubject<Map<number, Hotel>>(new Map<number, Hotel>)
  public offers = this.offerSubject.asObservable();

  private offerCart: Map<number, OfferCart> = new Map<number, OfferCart>();

  private currentSelectedOffer: Hotel | undefined;

  constructor(private readonly httpClient: HttpClient) { }

  public fetchOffers(dest: string | undefined, name: string | undefined, from: Date | undefined, to: Date | undefined, tags: Array<Tag> | undefined): void {
    this.offerSubject.value.clear();
    let params = new HttpParams();

    if (dest) {
      params = params.append('land', dest);
    }
    if (name) {
      params = params.append('name', name);
    }
    if (from && to) {
      params = params.append('from', from.toISOString());
      params = params.append('to', to.toISOString());
    }
    if (tags) {
      tags.forEach(tag => {
        params = params.append('tags', tag.id);
      });
    }

    this.httpClient.get(environment.HotelAPI + "hotels", { params: params }).subscribe((hotelsResponse) => {
      // Show only hits
      if (hotelsResponse || (hotelsResponse as Array<Hotel>).length == 0) {
        (hotelsResponse as Array<Hotel>).forEach((res: Hotel) => {
          if (res.travels.length > 0) {
            this.offerSubject.value.set(res.id, res)
          } else {
            this.offerSubject.value.delete(res.id)
          }
        })
        this.offerSubject.next(this.offerSubject.getValue());
      } else {
        this.offerSubject.next(new Map())
      }
    })

  }

  public selectOffer(id: number) {
    this.currentSelectedOffer = this.offerSubject.value.get(id);
  }

  public getSelectedOffer(): Hotel | undefined {
    return this.currentSelectedOffer;
  }

  public fetchOffersDemo() {
    this.offerSubject.value.clear();
    this.offerSubject.value.set(Hoteldemo.id, Hoteldemo)
    this.offerSubject.next(this.offerSubject.getValue());
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
    this.offerCart.set(offer.travelId, offer);
  }

  private handleError(error: any): Promise<any> {
    console.log('Interner Fehler: ' + error.message);
    return Promise.reject(error.message || error);
  }
}
