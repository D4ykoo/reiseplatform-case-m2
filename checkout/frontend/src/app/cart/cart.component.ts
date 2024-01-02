import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { environment } from '../../environment/environment';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-cart',
  standalone: true,
  imports: [CommonModule,RouterOutlet, RouterLink, RouterLinkActive],
  templateUrl: './cart.component.html',
  styleUrl: './cart.component.css',
})
export class CartComponent {
  hotelmanagementUrl = environment.hotelmanagement_url;

  cart = [
    { name: 'Hotel 1', price: 100 },
    { name: 'Hotel 2', price: 200 },
  ];

  constructor() {};

  public getTotal(): number {
    return this.cart.reduce((sum, item) => sum + item.price, 0);
  }

  public removeFromCart(index: number): void {
    this.cart.splice(index, 1);
  }

  public clearCart(): void {
    this.cart = [];
  }
}
