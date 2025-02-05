unit Unit1;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ComCtrls, StdCtrls,
  ButtonPanel, ExtCtrls, Buttons, ActnList, Menus, RTTICtrls, Types;

type

  { TFormMain }

  TFormMain = class(TForm)
    Action1: TAction;
    Action2: TAction;
    Action3: TAction;
    ActionList1: TActionList;
    Button_FangWenBaiDu: TButton;
    Button_GMM: TButton;
    Button_stop: TButton;
    Button_start: TButton;
    Edit1: TEdit;
    EdtThreads: TEdit;
    EdtName: TEdit;
    Image1: TImage;
    Label1: TLabel;
    Label2: TLabel;
    ListView1: TListView;
    Memo_log: TMemo;
    MenuItem1: TMenuItem;
    MenuItem_tuiChu: TMenuItem;
    MenuItem_FuZhiZhangHao: TMenuItem;
    MenuItem_FuZhiMiMa: TMenuItem;
    OpenDialog1: TOpenDialog;
    PageControl1: TPageControl;
    Panel1: TPanel;
    Panel2: TPanel;
    PopupMenu_tuPan: TPopupMenu;
    PopupMenu_list: TPopupMenu;
    ProgressBar1: TProgressBar;
    ScrollBox1: TScrollBox;
    StatusBar1: TStatusBar;
    TabSheet1: TTabSheet;
    TabSheet2: TTabSheet;
    TabSheet3: TTabSheet;
    TrayIcon1: TTrayIcon;
    procedure Action1Execute(Sender: TObject);
    procedure Button_FangWenBaiDuClick(Sender: TObject);
    procedure Button_GMMClick(Sender: TObject);
    procedure Button_stopClick(Sender: TObject);
    procedure Button_startClick(Sender: TObject);
    procedure FormCreate(Sender: TObject);
    procedure MenuItem1Click(Sender: TObject);
    procedure MenuItem_FuZhiZhangHaoClick(Sender: TObject);
    procedure MenuItem_FuZhiMiMaClick(Sender: TObject);
    procedure ScrollBox1Click(Sender: TObject);
    procedure TabSheet1ContextPopup(Sender: TObject; MousePos: TPoint;
      var Handled: Boolean);
    procedure ToolButton1Click(Sender: TObject);
  private

  public

  end;

var
  FormMain: TFormMain;

implementation

{$R *.lfm}

{ TFormMain }

procedure TFormMain.FormCreate(Sender: TObject);
begin

end;

procedure TFormMain.MenuItem1Click(Sender: TObject);
begin

end;

procedure TFormMain.MenuItem_FuZhiZhangHaoClick(Sender: TObject);
begin

end;

procedure TFormMain.MenuItem_FuZhiMiMaClick(Sender: TObject);
begin

end;

procedure TFormMain.ScrollBox1Click(Sender: TObject);
begin

end;



procedure TFormMain.TabSheet1ContextPopup(Sender: TObject; MousePos: TPoint;
  var Handled: Boolean);
begin

end;

procedure TFormMain.ToolButton1Click(Sender: TObject);
begin

end;

procedure TFormMain.Button_stopClick(Sender: TObject);
begin

end;

procedure TFormMain.Action1Execute(Sender: TObject);
begin

end;

procedure TFormMain.Button_FangWenBaiDuClick(Sender: TObject);
begin

end;

procedure TFormMain.Button_GMMClick(Sender: TObject);
begin

end;

procedure TFormMain.Button_startClick(Sender: TObject);
begin

end;

end.

