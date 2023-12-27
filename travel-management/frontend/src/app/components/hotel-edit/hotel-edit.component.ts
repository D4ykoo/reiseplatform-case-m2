import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { MultiSelectModule } from 'primeng/multiselect';
import { Tag } from 'primeng/tag';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../../environments/environment.development';
import { EditorModule } from 'primeng/editor';
import { InputMaskModule } from 'primeng/inputmask';
import { FileUploadModule } from 'primeng/fileupload';
import { CommonModule } from '@angular/common';
import { ContextMenuModule } from 'primeng/contextmenu';
import { CreateHotel, CreatePicture } from '../../models/requests';
import { catchError, lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-hotel-edit',
  standalone: true,
  imports: [MultiSelectModule, FormsModule, EditorModule, InputMaskModule, FileUploadModule, CommonModule, ContextMenuModule],
  templateUrl: './hotel-edit.component.html',
  styleUrl: './hotel-edit.component.css'
})

export class HotelEditComponent implements OnInit {

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

  constructor(private httpClient: HttpClient) {

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
    let createHotel: CreateHotel = {
      description: this.description as string, hotelname: this.hotelname,
      land: this.land, pictures: this.pictures, state: this.state, street: this.state,
      tagids: this.selectedTags, vendorid: 0, vendorname: "asdf"
    };
    lastValueFrom(this.httpClient.post(environment.HotelAPI + "hotels", createHotel)).then((e) => { this.clear(); console.log(e); }).catch((e) => console.log(e))
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
  ngOnInit() {

    this.httpClient.get(environment.HotelAPI + "tags").subscribe((res) => {
      if (res)
        this.tags = (res as Tag[]);
    })

    this.actions = [
      { label: 'Delete', icon: 'pi pi-fw pi-trash' }
    ];

  }

}

