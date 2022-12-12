import {Component, ElementRef, EventEmitter, OnInit, Output, ViewChild} from '@angular/core';
import {debounceTime, distinctUntilChanged, fromEvent, map} from "rxjs";

@Component({
  selector: 'app-searchbar',
  templateUrl: './searchbar.component.html',
  styleUrls: ['./searchbar.component.sass']
})
export class SearchbarComponent implements OnInit {

  @Output() onSearch = new EventEmitter<string>()

  @ViewChild('searchInput', {static: true}) searchInput!: ElementRef;
  searchTerm = '';

  constructor() { }

  ngOnInit(): void {
    fromEvent(this.searchInput.nativeElement, 'keyup').pipe(
        map((event: any) => event.target.value),
        debounceTime(500),
        distinctUntilChanged()
    ).subscribe((query: string) => {
      this.onSearch.emit(query)
    });
  }

}
