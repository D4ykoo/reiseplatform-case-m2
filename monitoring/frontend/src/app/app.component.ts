import { Component, OnDestroy, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { TabMenuModule } from 'primeng/tabmenu';
import { MenuItem } from 'primeng/api';
import { TableModule } from 'primeng/table';
import { Subscription, interval, startWith, switchMap } from 'rxjs';
import { UserService } from './service/user.service';
import { Events } from './model/response';

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
export class AppComponent implements OnInit, OnDestroy {
  items: MenuItem[] | undefined;
  timeInterval!: Subscription;
  userEvents!: Events[];
  travelEvents!: Events[];
  checkoutEvents!: Events[];
  activeItem: MenuItem | undefined;

  constructor(private userService: UserService) {}

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
    console.log(event);
    this.timeInterval?.unsubscribe();
    console.log('CHANGE ' + this.activeItem?.id);
    switch (this.activeItem?.id) {
      case '0':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.userService.getUserEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.userEvents = res.body as Events[];
              console.log(this.userEvents);
            } else {
              console.log('HTTP Error', res);
            }
          });
        break;
      case '1':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.userService.getCheckoutEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.checkoutEvents = res.body as Events[];
              console.log(this.checkoutEvents);
            } else {
              console.log('HTTP Error', res);
            }
          });
        break;
      case '2':
        this.timeInterval = interval(5000)
          .pipe(
            startWith(0),
            switchMap(() => this.userService.getTravelEvents()),
          )
          .subscribe((res) => {
            if (res.ok) {
              this.travelEvents = res.body as Events[];
              console.log(this.travelEvents);
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
