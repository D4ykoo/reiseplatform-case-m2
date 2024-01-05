import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { environment } from '../../environments/environment';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { CartService } from '../../services/cart.service';
import { CombinedCart } from '../../models/cart.models';

@Component({
  selector: 'app-cart',
  standalone: true,
  imports: [CommonModule, RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './cart.component.html',
  styleUrl: './cart.component.css',
})
export class CartComponent {
  constructor(private cartService: CartService) {}

  hotelmanagementUrl = environment.hotelmanagement_url;

  cartId = 1;



  cart: CombinedCart = {
    cart: undefined,
    hotel: []
  }

  ngOnInit() {
    this.getCart();
  }

  public getCart() {
    this.cartService.getCart(this.cartId).subscribe((res: any) => {
      this.cart = res;
      console.log("res: "+ res);
      console.log("cart: ", this.cart);
    });
  }

  public removeFromCart(hotelId: number, travelId: number): void {
    this.cart.hotel![hotelId].travels.splice(travelId, 1);

    this.cartService.removeFromCart(this.cart.cart!.id, hotelId, travelId).subscribe((res: any) => {
      console.log(res);
      // this.cart.splice(id, 1);
    });
  }

  public getTotal() {
    let total = 0;
    this.cart.hotel!.forEach((hotel) => {
      hotel.travels.forEach((travel) => {
        total += travel.price;
      });
    });
    return total;
  }

  public clearCart(): void {
    // this.cart = [];
  }
}
