import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartySidenavComponent } from './manage-party-sidenav.component';

describe('ManagePartySidenavComponent', () => {
  let component: ManagePartySidenavComponent;
  let fixture: ComponentFixture<ManagePartySidenavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartySidenavComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartySidenavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
