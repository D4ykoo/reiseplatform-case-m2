import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { environment } from '../../environments/environment';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { CartService } from '../../services/cart.service';

@Component({
  selector: 'app-cart',
  standalone: true,
  imports: [CommonModule,RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './cart.component.html',
  styleUrl: './cart.component.css',
})
export class CartComponent {
  constructor(
    private cartService: CartService,
  ) {}

  hotelmanagementUrl = environment.hotelmanagement_url;

  cartId = 1;

  cart = [
    { name: 'Hotel 1', price: 100 },
    { name: 'Hotel 2', price: 200 },
  ];


  ngOnInit() {
    this.getCart();
  }

  public getCart() {
    this.cartService.getCart(this.cartId).subscribe((res: any) => {
      for (const item of res) {
        this.cart.push(item);
      }
    });
  }

  public removeFromCart(id: number): void {
    this.cartService.removeFromCart(id).subscribe((res: any) => {
      console.log(res);
      this.cart.splice(id, 1);
    });
  }

  public addToCart(id: any) {
    this.cartService.addToCart(id);
  }


  public getTotal(): number {
    return this.cart.reduce((sum, item) => sum + item.price, 0);
  }



  public clearCart(): void {
    this.cart = [];
  }
}
