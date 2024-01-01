import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { MenuItem } from 'primeng/api';
import { TabMenuModule } from 'primeng/tabmenu';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule } from '@angular/forms';
import { HotelEditComponent } from "../hotel-edit/hotel-edit.component";
import { TravelOfferEditComponent } from "../travel-offer-edit/travel-offer-edit.component";
import { CommonModule } from '@angular/common';

interface TypeEdit {
  id: number;
  name: string;
}

@Component({
    selector: 'app-edit-panel',
    standalone: true,
    templateUrl: './edit-panel.component.html',
    styleUrl: './edit-panel.component.css',
    imports: [TabMenuModule, DropdownModule, FormsModule, HotelEditComponent, TravelOfferEditComponent, CommonModule]
})
export class EditPanelComponent implements OnInit {

  items!: MenuItem[];
  activeItem!: MenuItem;
  editType!: TypeEdit[] | undefined;
  showEditPanel: number = 0;
  selectedType: TypeEdit = { id: 0, name: "" };
  constructor(private readonly httpClient: HttpClient) {

  }

  ngOnInit() {
    this.items = [
      { label: 'New', icon: 'pi pi-fw pi-calendar' },
      { label: 'Edit', icon: 'pi pi-fw pi-pencil' },
    ];
    this.activeItem = this.items[0];
    this.editType = [{ id: 1, name: "Hotel" }, { id: 2, name: "Travel" }]
  }
}
