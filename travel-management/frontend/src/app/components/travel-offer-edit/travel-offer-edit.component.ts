import { CommonModule } from '@angular/common';
import { Component, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';
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
import { Travel } from '../../models/travel';

@Component({
  selector: 'app-travel-offer-edit',
  standalone: true,
  imports: [DropdownModule, CommonModule, FormsModule, InputNumberModule, CalendarModule, EditorModule],
  templateUrl: './travel-offer-edit.component.html',
  styleUrl: './travel-offer-edit.component.css'
})
export class TravelOfferEditComponent implements OnInit, OnChanges {

  @Input()
  editorMode!: string | undefined;

  hotels!: Hotel[];
  selectedHotel: Hotel | undefined;
  travels!: Travel[];
  selectedTravel!: Travel | undefined;;
  price!: number;
  public rangeDates: Date[] | undefined;
  description: string | undefined;

  constructor(private readonly httpClient: HttpClient) {
  }


  ngOnInit(): void {
    this.setup();
  }

  public setup() {
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
    if (this.editorMode == "New") {
      let createTravel: CreateTravel = {
        description: this.description as string, from: from.toISOString(), to: to.toISOString(), price: this.price, vendorid: 0, vendorname: ""
      };
      lastValueFrom(this.httpClient.post(environment.HotelAPI + "hotels/" + this.selectedHotel?.id + "/travels", createTravel)).then((e) => { this.clear(); console.log(e); }).catch((e) => console.log(e))
    }
    if (this.editorMode == "Edit") {
      let updateTravel: CreateTravel = {
        description: this.description as string, from: from.toISOString(), to: to.toISOString(), price: this.price, vendorid: 0, vendorname: ""
      };
      lastValueFrom(this.httpClient.put(environment.HotelAPI + "hotels/" + this.selectedHotel?.id + "/travels/" + this.selectedTravel?.id, updateTravel)).then(
        (res) => {
          if (res) {
            let tmp = (res as Travel)
            if (this.selectedTravel) {
              this.selectedTravel.id = tmp.id;
              this.selectedTravel.createdat = tmp.createdat;
              this.selectedTravel.description = tmp.description;
              this.selectedTravel.from = tmp.from;
              this.selectedTravel.price = tmp.price;
              this.selectedTravel.to = tmp.to;
              this.selectedTravel.updatedat = tmp.updatedat;
              this.selectedTravel.vendorid = tmp.vendorid;
              this.selectedTravel.vendorname = tmp.vendorname;
              this.loadSettings();
            } else {
              this.selectedTravel = tmp;
            }
          }
        }).catch((e) => console.log(e))
    }
  }

  loadSettings() {
    if (this.selectedTravel) {
      this.description = this.selectedTravel.description;
      this.price = this.selectedTravel.price;
      this.rangeDates = [new Date(this.selectedTravel.from), new Date(this.selectedTravel.from)];
    }
  }

  loadTravels() {
    if (this.editorMode == 'Edit' && this.selectedHotel) {
      this.travels = this.selectedHotel.travels
      this.selectedTravel = undefined;
    }
  }

  loadTraveldata() {
    if (this.editorMode == 'Edit' && this.selectedHotel && this.selectedTravel) {
      this.description = this.selectedTravel.description;
      this.price = this.selectedTravel.price;
      this.rangeDates = [new Date(this.selectedTravel.from), new Date(this.selectedTravel.to)];
    }

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
    this.travels = new Array()
    this.selectedTravel = undefined;
  }

  ngOnChanges(changes: SimpleChanges): void {
    const mode = changes['editorMode'];
    if (mode.currentValue != mode.previousValue && mode.currentValue == "Edit") {
      this.setup();
    } else {
      this.clear();
    }
  }
  delete() {
    if (this.selectedHotel && this.selectedTravel) {
      lastValueFrom(this.httpClient.delete(environment.HotelAPI + "hotels/" + this.selectedHotel?.id + "/travels/" + this.selectedTravel.id)).then(
        (e) => {
          console.log(e);
          this.clear()
        }).catch((e) => console.log(e));
    }
  }

}
