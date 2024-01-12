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

  hotelmanagementUrl = environment.hotelmanagementUrl;

  cartId = 1;

  cart: CombinedCart = {
    cart: undefined,
    hotel: [],
  };

  ngOnInit() {
    this.getCart();
  }

  public getCart() {
    this.cartService.getCart().subscribe((res) => {
      this.cart = res as CombinedCart;
    });
  }

  public removeFromCart(hotelId: number, travelId: number): void {
    this.cartService
      .removeFromCart(this.cart.cart!.id, hotelId, travelId)
      .subscribe(() => {
        this.cart.hotel![hotelId].travels.splice(travelId, 1);
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
