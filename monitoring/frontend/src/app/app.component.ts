import { Component, OnDestroy, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { TabMenuModule } from 'primeng/tabmenu';
import { MenuItem } from 'primeng/api';
import { TableModule } from 'primeng/table';
import { Subscription, interval, startWith, switchMap } from 'rxjs';
import { EventService } from './service/event.service';
import { Events } from './model/response';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

@Component({
  selector: 'app-root',
  standalone: true,
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
  imports: [
    CommonModule,
    RouterOutlet,
    TabMenuModule,
    TabMenuModule,
    TableModule,
    NavBarComponent,
  ],
})
export class AppComponent implements OnInit, OnDestroy {
  items: MenuItem[] | undefined;
  timeInterval!: Subscription;
  userEvents!: Events[];
  travelEvents!: Events[];
  checkoutEvents!: Events[];
  activeItem: MenuItem | undefined;

  constructor(private eventService: EventService) {}

  ngOnInit() {
    this.items = [
      { id: '0', label: 'User Logs', icon: 'pi pi-user' },
      { id: '1', label: 'Checkout Logs', icon: 'pi pi-shopping-cart' },
      { id: '2', label: 'Travel Logs', icon: 'pi pi-home' },
    ];
    this.activeItem = this.items[0];
    this.startpolling();
  }

  startpolling() {
    this.timeInterval?.unsubscribe();
    switch (this.activeItem?.id) {
      case '0':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.eventService.getUserEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.userEvents = res.body as Events[];
            } else {
              console.log('HTTP Error', res);
            }
          });
        break;
      case '1':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.eventService.getCheckoutEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.checkoutEvents = res.body as Events[];
            } else {
              console.log('HTTP Error', res);
            }
          });
        break;
      case '2':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.eventService.getTravelEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.travelEvents = res.body as Events[];
            } else {
              console.log('HTTP Error', res);
            }
          });
        break;
    }
  }

  ngOnDestroy(): void {
    this.timeInterval.unsubscribe();
  }
}
