import { Component } from '@angular/core';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';

@Component({
  selector: 'app-info-card',
  standalone: true,
  imports: [CardModule,ButtonModule, DialogModule],
  templateUrl: './info-card.component.html',
  styleUrl: './info-card.component.css'
})
export class InfoCardComponent {
  visible: boolean = false;

  showDialog() {
      this.visible = !this.visible;
  }
}
