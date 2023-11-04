import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {CommonModule} from "@angular/common";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {MatButtonModule} from "@angular/material/button";

@Component({
  standalone: true,
  imports: [CommonModule, MatProgressSpinnerModule, MatButtonModule],
  selector: 'cheers-button',
  templateUrl: './cheers-button.component.html',
  styleUrls: ['./cheers-button.component.sass']
})
export class CheersButtonComponent implements OnInit {

  @Input() isLoading: boolean = false
  @Input() text: string = ""
  @Input() disabled: boolean = false

  constructor() { }

  ngOnInit(): void {
  }

}
