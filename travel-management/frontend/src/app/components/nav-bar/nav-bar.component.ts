import { Component } from '@angular/core';
import { Router } from '@angular/router';
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

navigateToStatMngt() {
}
navigateToUserMngt() {
throw new Error('Method not implemented.');
}
navigateToTravelMngt() {
  this.router.navigate(['/management']);
}
navigateToSearch() {
  this.router.navigate(['/']);
}
  
  sidebarVisible: boolean = false;
  
  constructor(private readonly router: Router){
    
  }

  public onNavigate(){
    window.open("https://www.google.com?cat=55","_self");
}
}
