import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
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
