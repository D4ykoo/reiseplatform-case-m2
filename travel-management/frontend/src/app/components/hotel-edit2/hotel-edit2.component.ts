import { Component, Input, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MultiSelectModule } from 'primeng/multiselect';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { EditorModule } from 'primeng/editor';
import { InputMaskModule } from 'primeng/inputmask';
import { FileUploadModule } from 'primeng/fileupload';
import { CommonModule } from '@angular/common';
import { ContextMenuModule } from 'primeng/contextmenu';
import { CreateHotel, CreatePicture } from '../../models/requests';
import { Hotel } from '../../models/hotel';
import { DropdownModule } from 'primeng/dropdown';
import { lastValueFrom } from 'rxjs';
import { Tag } from '../../models/tag';

@Component({
  selector: 'app-hotel-edit2',
  standalone: true,
  imports: [MultiSelectModule, FormsModule, EditorModule, InputMaskModule, FileUploadModule, CommonModule, ContextMenuModule, DropdownModule],
  templateUrl: './hotel-edit2.component.html',
  styleUrl: './hotel-edit2.component.css'
})


export class HotelEdit2Component implements OnInit {

  constructor(private readonly httpClient: HttpClient) {

  }


  hotels!: Hotel[];
  hotel!: Hotel;
  tags!: Tag[];
  selectedTags!: Tag[];
  description: string | undefined;
  images = new Array<any>();
  pictures = new Array<CreatePicture>();
  actions: MenuItem[] | undefined;
  hotelname!: string;
  street!: string;
  state!: string;
  land!: string;

  ngOnInit(): void {
    lastValueFrom(this.httpClient.get(environment.HotelAPI + "hotels")).then((res) => {
      if (res)
        this.hotels = (res as Hotel[]);
    })
    lastValueFrom(this.httpClient.get(environment.HotelAPI + "tags")).then((res) => {
      if (res)
        this.tags = (res as Tag[]);
    })

  }

  loadSettings() {
    this.hotelname = this.hotel.hotelname;
    this.description = this.hotel.description;
    this.land = this.hotel.land;
    this.state = this.hotel.state;
    this.street = this.hotel.street;
    this.selectedTags = this.hotel.tags;
    this.images = this.hotel.pictures
    this.pictures = this.pictures;
  }

  onFileChange(event: any) {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];

      const reader = new FileReader();
      reader.onload = e => {
        this.images.push(reader.result);
        this.pictures.push({ id: 0, description: "", payload: reader.result as string });
      };

      reader.readAsDataURL(file);
    }
  }

  delete(index: number) {
    if (index > -1) {
      this.images.splice(index, 1);
      this.pictures.splice(index, 1);
    }
  }

  submit() {
    let UpdateHotel: Hotel = {
      description: this.description as string, hotelname: this.hotelname,
      land: this.land, pictures: this.pictures, state: this.state, street: this.state,
      tags: this.selectedTags, vendorid: 0, vendorname: "asdf", id: this.hotel.id, travels: this.hotel.travels
    };
    lastValueFrom(this.httpClient.put(environment.HotelAPI + "hotels", UpdateHotel)).then((e) => { this.clear(); console.log(e); }).catch((e) => console.log(e))
  }


  clear() {
    lastValueFrom(this.httpClient.get(environment.HotelAPI + "tags")).then(res => {
      if (res)
        this.tags = (res as Tag[]);
    })
    this.selectedTags = new Array();
    this.description = "";
    this.images = new Array<any>();
    this.pictures = new Array<CreatePicture>();
    this.hotelname = "";
    this.street = "";
    this.state = "";
    this.land = "";

  }
}
