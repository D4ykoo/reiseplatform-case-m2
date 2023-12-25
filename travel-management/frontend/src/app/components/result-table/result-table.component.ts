import { Component, OnDestroy, OnInit } from '@angular/core';
import { DataViewModule } from 'primeng/dataview';
import { FormsModule } from '@angular/forms';
import { TagModule } from 'primeng/tag';
import { CommonModule } from '@angular/common';
import { Hotel } from '../../models/hotel';
import { ButtonModule } from 'primeng/button';
import { OfferService } from '../../services/offer.service';

@Component({
  selector: 'app-result-table',
  standalone: true,
  imports: [DataViewModule, FormsModule, TagModule, CommonModule, ButtonModule],
  templateUrl: './result-table.component.html',
  styleUrl: './result-table.component.css'
})
export class ResultTableComponent implements OnInit {
  public hotels!: Hotel[];

  constructor(private readonly offerService: OfferService) {

  }

  ngOnInit(): void {

    this.offerService.offers.subscribe((a)=>{
      this.hotels = Array.from(a.values());
    })
  }

}
