import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'admin-button',
  templateUrl: './admin-button.component.html',
  styleUrls: ['./admin-button.component.sass']
})
export class AdminButtonComponent implements OnInit {

  @Input() isAdmin: boolean;
  // @Output() click = new EventEmitter();

  constructor() { }

  ngOnInit(): void {
  }

}
