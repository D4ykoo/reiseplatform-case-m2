import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { DropdownModule } from 'primeng/dropdown';
import { Hotel } from '../../models/hotel';
import { FormsModule } from '@angular/forms';
import { InputNumberModule } from 'primeng/inputnumber';
import { CalendarModule } from 'primeng/calendar';
import { EditorModule } from 'primeng/editor';
import { lastValueFrom } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { CreateTravel } from '../../models/requests';

@Component({
  selector: 'app-travel-offer-edit',
  standalone: true,
  imports: [DropdownModule, CommonModule, FormsModule, InputNumberModule, CalendarModule, EditorModule],
  templateUrl: './travel-offer-edit.component.html',
  styleUrl: './travel-offer-edit.component.css'
})
export class TravelOfferEditComponent implements OnInit {

  hotels!: Hotel[] | undefined;
  selectedHotel: Hotel | undefined;
  price!: number;
  public rangeDates: Date[] | undefined;
  description: string | undefined;

  constructor(private readonly httpClient: HttpClient) {
  }

  ngOnInit(): void {
    lastValueFrom(this.httpClient.get(environment.HotelAPI + "hotels")).then((hotels) => {
      if (hotels)
        this.hotels = (hotels as Hotel[])
    })
  }

  submit() {
    let from = new Date();
    let to = new Date();

    if (this.rangeDates) {
      from = this.rangeDates[0];
      to = this.rangeDates[1];
      if (!to) {
        to = from
      }
    }

    let createTravel: CreateTravel = {

      description: this.description as string, from: from.toISOString(), to: to.toISOString(), price: this.price, vendorid: 0, vendorname: "asdf"
    };
    lastValueFrom(this.httpClient.post(environment.HotelAPI + "hotels/"+ this.selectedHotel?.id+"/travels", createTravel)).then((e) => { this.clear(); console.log(e); }).catch((e) => console.log(e))
  }


  clear() {
    lastValueFrom(this.httpClient.get(environment.HotelAPI + "hotels")).then(res => {
      if (res)
        this.hotels = (res as Hotel[]);
    })
    this.selectedHotel = undefined;
    this.description = "";
    this.price = 0;
    this.rangeDates = new Array();
  }
}
