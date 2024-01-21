import { Component, OnDestroy, OnInit } from '@angular/core';
import { DataViewModule } from 'primeng/dataview';
import { FormsModule } from '@angular/forms';
import { TagModule } from 'primeng/tag';
import { CommonModule } from '@angular/common';
import { Hotel } from '../../models/hotel';
import { ButtonModule } from 'primeng/button';
import { OfferService } from '../../services/offer.service';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { InfoCardComponent } from '../info-card/info-card.component';

@Component({
  selector: 'app-result-table',
  standalone: true,
  imports: [DataViewModule, FormsModule, TagModule, CommonModule, ButtonModule],
  templateUrl: './result-table.component.html',
  styleUrl: './result-table.component.css',
  providers: [DialogService],
})
export class ResultTableComponent implements OnInit, OnDestroy {
  public hotels!: Hotel[];

  ref: DynamicDialogRef | undefined;

  constructor(
    private readonly offerService: OfferService,
    public dialogService: DialogService,
  ) {}

  public detail(id: number) {
    this.offerService.selectOffer(id);
    this.ref = this.dialogService.open(InfoCardComponent, {
      width: '50vw',
      modal: true,
      breakpoints: {
        '2000px': '65vw',
        '1200px': '80vw',
        '640px': '95vw',
      },
    });
  }
  ngOnInit(): void {
    this.offerService.offers.subscribe((a) => {
      this.hotels = Array.from(a.values());
    });
  }

  ngOnDestroy() {
    if (this.ref) {
      this.ref.close();
    }
  }
}
