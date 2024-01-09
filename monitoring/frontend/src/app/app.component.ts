import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { TabMenuModule } from 'primeng/tabmenu';
import { MenuItem } from 'primeng/api';
import { TableModule } from 'primeng/table';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    TabMenuModule,
    TabMenuModule,
    TableModule,
  ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent implements OnInit {
  products!: any[];
  items: MenuItem[] | undefined;

  ngOnInit() {
    this.items = [
      { label: 'User Logs', icon: 'pi pi-user' },
      { label: 'Checkout Logs', icon: 'pi pi-shopping-cart' },
      { label: 'Travel Logs', icon: 'pi pi-home' },
    ];
  }
}
