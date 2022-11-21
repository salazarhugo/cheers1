import { Injectable } from '@angular/core';
import {BehaviorSubject, firstValueFrom, Observable, Subject} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class ThemeService {
  private _darkTheme = new BehaviorSubject<boolean>(true)
  isDarkTheme: Observable<boolean> = this._darkTheme.asObservable()

  setDarkTheme(isDarkTheme: boolean): void {
    this._darkTheme.next(isDarkTheme)
  }

  constructor() { }
}
