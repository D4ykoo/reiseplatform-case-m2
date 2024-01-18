import { Component, OnInit } from '@angular/core';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { OfferService } from '../../services/offer.service';
import { Hotel } from '../../models/hotel';
import { CommonModule } from '@angular/common';
import { GalleriaModule } from 'primeng/galleria';
import { FormsModule } from '@angular/forms';
import { DataViewModule } from 'primeng/dataview';
import { TagModule } from 'primeng/tag';
import { DynamicDialogRef } from 'primeng/dynamicdialog';
import { LoginService } from '../../services/login.service';
import { User } from '../../models/user';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-info-card',
  standalone: true,
  imports: [
    CardModule,
    ButtonModule,
    DialogModule,
    CommonModule,
    GalleriaModule,
    FormsModule,
    DataViewModule,
    TagModule,
  ],
  templateUrl: './info-card.component.html',
  styleUrl: './info-card.component.css',
})
export class InfoCardComponent implements OnInit {
  hotelOffer: Hotel | undefined;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  responsiveOptions: any[] | undefined;

  user: User | undefined;
  constructor(
    private offerService: OfferService,
    private httpClient: HttpClient,
    private ref: DynamicDialogRef,
    private loginService: LoginService,
  ) { }

  ngOnInit(): void {
    this.hotelOffer = this.offerService.getSelectedOffer();
    this.loginService.user.subscribe((u) => this.user = u);
    this.responsiveOptions = [
      {
        breakpoint: '1024px',
        numVisible: 5,
      },
      {
        breakpoint: '768px',
        numVisible: 3,
      },
      {
        breakpoint: '560px',
        numVisible: 1,
      },
    ];
  }

  closeDialog() {
    this.ref.close();
  }

  addToCart(id: number) {
    lastValueFrom(this.httpClient.put(
      environment.Checkout_API +
      'cart/addtocart/' + (this.user?.id as number) + '/' +
      (this.hotelOffer as Hotel).id +
      '/' +
      (this.hotelOffer as Hotel).travels[id].id, {}, { withCredentials: true }
    ));
    this.closeDialog();
  }
}
