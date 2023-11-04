import {Component, Input, OnInit } from '@angular/core';
import {CommonModule} from "@angular/common";
import {MatButtonModule} from "@angular/material/button";

@Component({
  standalone: true,
  selector: 'admin-button',
  imports: [
    CommonModule,
    MatButtonModule,
  ],
  templateUrl: './admin-button.component.html',
  styleUrls: ['./admin-button.component.sass'],
})
export class AdminButtonComponent  {

  @Input() isAdmin: boolean;
  // @Output() click = new EventEmitter();

  constructor() { }

}
