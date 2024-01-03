import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  Router,
  RouterLink,
  RouterLinkActive,
  RouterOutlet,
} from '@angular/router';

@Component({
  selector: 'app-payment',
  standalone: true,
  imports: [CommonModule, RouterLink, RouterLinkActive, RouterOutlet],
  templateUrl: './payment.component.html',
  styleUrl: './payment.component.css',
})
export class PaymentComponent {
  loading = false;
  showToast = false;
  constructor(private router: Router) {}

  private wait(ms: number): Promise<void> {
    return new Promise<void>((resolve) => setTimeout(resolve, ms));
  }

  public async pay() {
    console.log('Payment successful!');
    this.loading = true;

    await this.wait(3000);

    this.loading = false;

    this.showToast = true;

    await this.wait(3000);

    this.showToast = false;

    this.router.navigate(['/']); // TODO: Navigate to the hotel management page
  }
}
