import { Component, OnInit } from '@angular/core';
import { TabMenuModule } from 'primeng/tabmenu';
import { MenuItem } from 'primeng/api';

@Component({
  selector: 'app-edit-panel',
  standalone: true,
  imports: [TabMenuModule],
  templateUrl: './edit-panel.component.html',
  styleUrl: './edit-panel.component.css'
})
export class EditPanelComponent implements OnInit{


  items: MenuItem[] | undefined;
  activeItem: MenuItem | undefined;

  ngOnInit() {
      this.items = [
          { label: 'Home', icon: 'pi pi-fw pi-home' },
          { label: 'New', icon: 'pi pi-fw pi-calendar' },
          { label: 'Edit', icon: 'pi pi-fw pi-pencil' },
      ];
      this.activeItem = this.items[0];
  }

  changeView(event:any) {
  console.log(event);
  }
}
