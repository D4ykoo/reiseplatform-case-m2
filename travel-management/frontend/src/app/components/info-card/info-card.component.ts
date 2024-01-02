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
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';

@Component({
  selector: 'app-info-card',
  standalone: true,
  imports: [CardModule, ButtonModule, DialogModule, CommonModule, GalleriaModule, FormsModule,DataViewModule,TagModule],
  templateUrl: './info-card.component.html',
  styleUrl: './info-card.component.css'
})
export class InfoCardComponent implements OnInit {

  hotelOffer: Hotel | undefined;
  responsiveOptions: any[] | undefined;
  constructor(private offerService: OfferService,  private dialogService: DialogService, private ref: DynamicDialogRef){
  }

  ngOnInit(): void {
    this.hotelOffer = this.offerService.getSelectedOffer();
    this.responsiveOptions = [
      {
          breakpoint: '1024px',
          numVisible: 5
      },
      {
          breakpoint: '768px',
          numVisible: 3
      },
      {
          breakpoint: '560px',
          numVisible: 1
      }
  ];
  }

closeDialog() {
    this.ref.close();
}

buy(id: number) {
  this.offerService.addToCart({hotelId: (this.hotelOffer as Hotel).id,travelId:(this.hotelOffer as Hotel).travels[id].id, userId: 0 });
  this.closeDialog()
}
}
