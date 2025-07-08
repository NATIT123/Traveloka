import { Component } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'auth-dialog',
  templateUrl: 'auth.component.html',
  //   styleUrl: 'header.component.css',
  imports: [DialogModule, ButtonModule],
})
export class AuthComponent {
  visible: boolean = false;

  showDialog() {
    this.visible = true;
  }
}
