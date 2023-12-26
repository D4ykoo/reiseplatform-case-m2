import { Component, OnInit } from '@angular/core';
import { TabMenuModule } from 'primeng/tabmenu';
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

@Component({
  selector: 'app-edit-panel',
  standalone: true,
  imports: [TabMenuModule, MultiSelectModule, FormsModule, EditorModule, InputMaskModule, FileUploadModule, CommonModule, ContextMenuModule],
  templateUrl: './edit-panel.component.html',
  styleUrl: './edit-panel.component.css'
})
export class EditPanelComponent implements OnInit {
  tags!: Tag[];
  selectedTags!: Tag[];
  description: string | undefined;
  images = new Array<any>();
  actions: MenuItem[] | undefined;
  items: MenuItem[] | undefined;
  activeItem: MenuItem | undefined;

  constructor(private readonly httpClient: HttpClient) {

  }

  onFileChange(event: any) {
    if (event.target.files && event.target.files[0]) {
      const file = event.target.files[0];

      const reader = new FileReader();
      reader.onload = e => this.images.push(reader.result);

      reader.readAsDataURL(file);
    }
  }
  delete(index: number) {
    if (index > -1) {
      this.images.splice(index, 1);
    }
  }

  ngOnInit() {
    this.items = [
      { label: 'Home', icon: 'pi pi-fw pi-home' },
      { label: 'New', icon: 'pi pi-fw pi-calendar' },
      { label: 'Edit', icon: 'pi pi-fw pi-pencil' },
    ];
    this.activeItem = this.items[0];
    this.httpClient.get(environment.HotelAPI + "tags").subscribe((res) => {
      if (res)
        this.tags = (res as Tag[]);
    })

    this.actions = [
      { label: 'Delete', icon: 'pi pi-fw pi-trash' }
    ];

  }

  changeView(event: any) {
    console.log(event);
  }
}
