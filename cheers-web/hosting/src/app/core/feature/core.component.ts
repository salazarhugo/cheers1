import {AfterViewInit, Component, OnInit } from '@angular/core';
import {UserService} from "../../shared/data/services/user.service";
import {fadeAnimation} from "../../animtaions";
import {MatDrawerMode} from "@angular/material/sidenav";
import {BreakpointObserver, Breakpoints} from "@angular/cdk/layout";
import {distinctUntilChanged, tap} from "rxjs";

@Component({
    selector: 'app-core',
    templateUrl: './core.component.html',
    styleUrls: ['./core.component.sass'],
    animations: [fadeAnimation],
})
export class CoreComponent implements OnInit, AfterViewInit {

    mode: MatDrawerMode = "side"

    readonly breakpoint$ = this.breakpointObserver
        .observe([Breakpoints.Large, Breakpoints.Medium, Breakpoints.Small, '(min-width: 500px)'])
        .pipe(
            tap(value => console.log(value)),
            distinctUntilChanged()
        );

    constructor(
        public userService: UserService,
        private breakpointObserver: BreakpointObserver,
    ) {
    }

    ngOnInit(): void {
        this.breakpoint$.subscribe(() =>
            this.breakpointChanged()
        );
    }

    private breakpointChanged() {
        if(this.breakpointObserver.isMatched(Breakpoints.Large)) {
            this.mode = "side"
        } else if(this.breakpointObserver.isMatched(Breakpoints.Medium)) {
            this.mode = "side"
        } else if(this.breakpointObserver.isMatched(Breakpoints.Small)) {
            this.mode = "over"
        }
    }

    ngAfterViewInit(): void {
    }
}
