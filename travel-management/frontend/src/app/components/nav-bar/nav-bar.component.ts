import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button';
import { SidebarModule } from 'primeng/sidebar';

@Component({
  selector: 'app-nav-bar',
  standalone: true,
  imports: [ButtonModule, SidebarModule],
  templateUrl: './nav-bar.component.html',
  styleUrl: './nav-bar.component.css'
})
export class NavBarComponent {
  
  sidebarVisible: boolean = false;
  
  public onNavigate(){
    window.open("https://www.google.com","_self");
}
}
