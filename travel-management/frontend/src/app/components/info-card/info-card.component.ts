import { Component, OnInit } from '@angular/core';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { OfferService } from '../../services/offer.service';

@Component({
  selector: 'app-info-card',
  standalone: true,
  imports: [CardModule,ButtonModule, DialogModule],
  templateUrl: './info-card.component.html',
  styleUrl: './info-card.component.css'
})
export class InfoCardComponent implements OnInit{
  constructor(private offerService: OfferService){

  }
  ngOnInit(): void {
    this.offerService.offers.subscribe((a)=>{
      console.log(a);
    })
  }

  show(){
    console.log("press button")
    this.offerService.fetchOffers();
  }
/*
  visible: boolean = false;

  showDialog() {
      this.visible = !this.visible;
  }
  */
}
