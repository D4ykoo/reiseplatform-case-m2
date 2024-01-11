import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SidebarModule } from 'primeng/sidebar';
import { environment } from '../../../environments/environment.development';
import { LoginService } from '../../services/login.service';
import { User } from '../../models/user';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-nav-bar',
  standalone: true,
  imports: [ButtonModule, SidebarModule, CommonModule, FormsModule],
  templateUrl: './nav-bar.component.html',
  styleUrl: './nav-bar.component.css',
})
export class NavBarComponent implements OnInit {
  sidebarVisible: boolean = false;

  user!: User;
  constructor(
    private readonly router: Router,
    private loginService: LoginService,
  ) {}

  ngOnInit(): void {
    this.loginService.user.subscribe((loginStatus) => {
      this.user = loginStatus;
    });
  }

  navigateToMonitoring() {
    const url = environment.Monitor_URL as unknown as string;
    window.open(url, '_self');
  }
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

  public onNavigate() {
    const url = environment.Login_URL as unknown as string;
    window.open(url, '_self');
  }
}
