import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SidebarModule } from 'primeng/sidebar';
import { environment } from '../../../environments/environment.development';

@Component({
  selector: 'app-nav-bar',
  standalone: true,
  imports: [ButtonModule, SidebarModule],
  templateUrl: './nav-bar.component.html',
  styleUrl: './nav-bar.component.css',
})
export class NavBarComponent {
  navigateToStatMngt() {}
  navigateToUserMngt() {
    const url = environment.Login_URL as unknown as string;
    window.open(url, '_self');
  }
  navigateToTravelMngt() {
    this.router.navigate(['/management']);
  }
  navigateToSearch() {
    this.router.navigate(['/']);
  }

  sidebarVisible: boolean = false;

  constructor(private readonly router: Router) {}

  public onNavigate() {
    const url = environment.Login_URL as unknown as string;
    window.open(url, '_self');
  }
}
