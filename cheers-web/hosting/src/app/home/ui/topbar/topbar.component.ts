import {Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {MatSidenav} from "@angular/material/sidenav";
import {debounceTime, distinctUntilChanged, fromEvent, map, Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {AuthService} from "../../../shared/data/services/auth.service";
import {Router} from "@angular/router";
import {ThemeService} from "../../../core/data/theme.service";
import {UserItem} from "../../../../gen/ts/cheers/type/user/user";

@Component({
    selector: 'app-topbar',
    templateUrl: './topbar.component.html',
    styleUrls: ['./topbar.component.sass']
})
export class TopbarComponent implements OnInit {

    isDarkTheme: Observable<boolean>
    $user: Observable<User | null> = of(null)
    $searchResults: Observable<User[] | null> = of(null)
    $recentSearches: Observable<User[] | null> = of(null)

    @ViewChild(MatSidenav) sidenav!: MatSidenav;
    @ViewChild('searchInput', {static: true}) searchInput!: ElementRef;

    constructor(
        public userService: UserService,
        private authService: AuthService,
        public themeService: ThemeService,
        private router: Router,
    ) {
        this.$user = this.userService.user$

    }

    ngOnInit(): void {
        this.isDarkTheme = this.themeService.isDarkTheme
        fromEvent(this.searchInput.nativeElement, 'keyup').pipe(
            map((event: any) => event.target.value),
            debounceTime(500),
            distinctUntilChanged()
        ).subscribe((query: string) => {
            this.searchUser(query)
        });
    }

    searchUser(query: string) {
        this.$searchResults = this.userService.searchUser(query)
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    async onSignOut() {
        await this.authService.signOut()
        await this.router.navigate(['sign-in'])
    }

    onUserClick(user: User) {
        localStorage.setItem('recent', JSON.stringify(user));
        this.$recentSearches = of([user])
    }

}
