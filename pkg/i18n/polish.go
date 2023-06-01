package i18n

func polishTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "Za mało miejsca do wyświetlenia paneli",
		DiffTitle:                           "Różnice",
		FilesTitle:                          "Pliki",
		BranchesTitle:                       "Gałęzie",
		CommitsTitle:                        "Commity",
		StashTitle:                          "Schowek",
		UnstagedChanges:                     "Zmiany poza poczekalnią",
		StagedChanges:                       "Zmiany w poczekalni",
		CommitSummary:                       "Komunikat commita",
		CredentialsUsername:                 "Użytkownik",
		CredentialsPassword:                 "Hasło",
		CredentialsPassphrase:               "Fraza",
		PassUnameWrong:                      "Niewłaściwe hasło, fraza lub nazwa użytkownika",
		CommitChanges:                       "Zatwierdź zmiany",
		AmendLastCommit:                     "Zmień ostatni commit",
		AmendLastCommitTitle:                "Zmień ostatni commit",
		SureToAmend:                         "Czy na pewno chcesz zmienić ostatni commit? Możesz zmienić komunikat commitu z panelu commitów.",
		NoCommitToAmend:                     "Brak commitów do zmiany.",
		CommitChangesWithEditor:             "Zatwierdź zmiany używając edytora",
		StatusTitle:                         "Status",
		GlobalTitle:                         "Globalne",
		Menu:                                "Menu",
		Execute:                             "Wykonaj",
		ToggleStaged:                        "Przełącz stan poczekalni",
		ToggleStagedAll:                     "Przełącz stan poczekalni wszystkich",
		Refresh:                             "Odśwież",
		Scroll:                              "Przewiń",
		FilterStagedFiles:                   "Pokaż tylko pliki w poczekalni",
		FilterUnstagedFiles:                 "Pokaż tylko pliki poza poczekalnią",
		ResetCommitFilterState:              "Resetuj filtr commitów",
		Checkout:                            "Przełącz",
		NoChangedFiles:                      "Brak zmienionych plików",
		PullWait:                            "Pobieranie zmian...",
		PushWait:                            "Wysyłanie zmian...",
		FetchWait:                           "Pobieram...",
		AlreadyCheckedOutBranch:             "Już przęłączono na tą gałąź",
		SureForceCheckout:                   "Jesteś pewny, że chcesz wymusić przełączenie? Stracisz wszystkie lokalne zmiany",
		ForceCheckoutBranch:                 "Wymuś przełączenie gałęzi",
		BranchName:                          "Nazwa gałęzi",
		NewBranchNameBranchOff:              "Nazwa nowej gałęzi (gałąź na bazie '{{.branchName}}')",
		CantDeleteCheckOutBranch:            "Nie możesz usunąć obecnie przełączonej gałęzi!",
		DeleteBranch:                        "Usuń gałąź",
		DeleteBranchMessage:                 "Jesteś pewien, że chcesz usunąć gałąź '{{.selectedBranchName}}' ?",
		ForceDeleteBranchMessage:            "Na pewno wymusić usunięcie gałęzi '{{.selectedBranchName}}'?",
		RebaseBranch:                        "Zmiana bazy gałęzi",
		CantRebaseOntoSelf:                  "Nie możesz zmienić bazy gałęzi na nią samą",
		CantMergeBranchIntoItself:           "Nie możesz scalić gałęzi do samej siebie",
		ForceCheckout:                       "Wymuś przełączenie",
		CheckoutByName:                      "Przełącz używając nazwy",
		NewBranch:                           "Nowa gałąź",
		NoBranchesThisRepo:                  "Brak gałęzi dla tego repozytorium",
		CommitWithoutMessageErr:             "Nie możesz commitować bez komunikatu",
		CloseCancel:                         "Zamknij",
		Confirm:                             "Potwierdź",
		Close:                               "Zamknij",
		SquashDown:                          "Ściśnij",
		FixupCommit:                         "Napraw commit",
		NoCommitsThisBranch:                 "Brak commitów dla tej gałęzi",
		CannotSquashOrFixupFirstCommit:      "There's no commit below to squash into",
		Fixup:                               "Napraw",
		SureFixupThisCommit:                 "Jesteś pewny, ze chcesz naprawić ten commit? Commit poniżej zostanie spłaszczony w górę wraz z tym",
		RewordCommit:                        "Zmień nazwę commita",
		RenameCommitEditor:                  "Zmień nazwę commita w edytorze",
		Error:                               "Błąd",
		PickHunk:                            "Wybierz kawałek",
		PickAllHunks:                        "Wybierz oba kawałki",
		Undo:                                "Cofnij",
		Pop:                                 "Wyciągnij",
		Drop:                                "Porzuć",
		Apply:                               "Zastosuj",
		NoStashEntries:                      "Brak pozycji w schowku",
		StashDrop:                           "Porzuć schowek",
		SureDropStashEntry:                  "Jesteś pewny, że chcesz porzucić tę pozycję w schowku?",
		NoTrackedStagedFilesStash:           "Nie masz śledzonych/zatwierdzonych plików do przechowania",
		StashChanges:                        "Przechowaj zmiany",
		RenameStash:                         "Rename stash",
		RenameStashPrompt:                   "Rename stash: {{.stashName}}",
		OpenConfig:                          "Otwórz konfigurację",
		EditConfig:                          "Edytuj konfigurację",
		ForcePush:                           "Wymuś wysłanie",
		ForcePushPrompt:                     "Twoja gałąź rozeszła się z gałęzią zdalną. Wciśnij 'esc' aby anulować lub 'enter' aby wymusić wysłanie.",
		ForcePushDisabled:                   "Twoja gałąź rozeszła się z gałęzią zdalną i wyłączyłeś wymuszenie wysłania",
		UpdatesRejectedAndForcePushDisabled: "Aktualizacje zostały odrzucone i wyłączyłeś wymuszenie wysłania",
		CheckForUpdate:                      "Sprawdź aktualizacje",
		CheckingForUpdates:                  "Sprawdzanie aktualizacji...",
		OnLatestVersionErr:                  "Już posiadasz najnowszą wersję",
		MajorVersionErr:                     "Nowa wersja ({{.newVersion}}) posiada niekompatybilne zmiany w porównaniu do obecnej wersji ({{.currentVersion}})",
		CouldNotFindBinaryErr:               "Nie można znaleźć pliku binarnego w {{.url}}",
		EditFile:                            "Edytuj plik",
		OpenFile:                            "Otwórz plik",
		IgnoreFile:                          "Dodaj do .gitignore",
		RefreshFiles:                        "Odśwież pliki",
		MergeIntoCurrentBranch:              "Scal do obecnej gałęzi",
		ConfirmQuit:                         "Na pewno chcesz wyjść z programu?",
		AllBranchesLogGraph:                 "Pokaż wszystkie logi gałęzi",
		UnsupportedGitService:               "Nieobsługiwana usługa git",
		CreatePullRequest:                   "Utwórz żądanie pobrania",
		CopyPullRequestURL:                  "Skopiuj adres URL żądania pobrania do schowka",
		NoBranchOnRemote:                    "Ta gałąź nie istnieje w zdalnym repo. Najpierw musisz ją wysłać.",
		Fetch:                               "Pobierz",
		NoAutomaticGitFetchTitle:            `Brak automatycznego "git fetch"`,
		NoAutomaticGitFetchBody:             `Lazygit nie może użyć "git fetch" w prywatnym repo, użyj f w panelu gałęzi żeby uruchomić "git fetch" ręcznie`,
		FileEnter:                           "Zatwierdź pojedyncze linie",
		FileStagingRequirements:             "Można tylko zatwierdzić pojedyncze linie dla śledzonych plików z niezatwierdzonymi zmianami",
		StagingTitle:                        "Poczekalnia",
		ReturnToFilesPanel:                  "Wróć do panelu plików",
		MergingTitle:                        "Scalanie",
		ConfirmMerge:                        "Czy na pewno chcesz scalić '{{.selectedBranch}}' do '{{.checkedOutBranch}}'?",
		FwdNoUpstream:                       "Nie można przewinąć gałęzi bez gałęzi nadrzędnej",
		FwdCommitsToPush:                    "Nie można przewinąć gałęzi z commitami do wysłania",
		ErrorOccurred:                       "Wystąpił błąd! Zgłoś problem na",
		MainTitle:                           "Główne",
		NormalTitle:                         "Zwykłe",
		SoftReset:                           "Miękki reset",
		SureSquashThisCommit:                "Czy na pewno spłaszczyć ten commit do commita niżej?",
		Squash:                              "Spłaszcz",
		PickCommit:                          "Wybierz commit (podczas zmiany bazy)",
		RevertCommit:                        "Odwróć commit",
		DeleteCommit:                        "Usuń commit",
		MoveDownCommit:                      "Przenieś commit 1 w dół",
		MoveUpCommit:                        "Przenieś commit 1 w górę",
		EditCommit:                          "Edytuj commit",
		AmendToCommit:                       "Popraw commit zmianami z poczekalni",
		FoundConflictsTitle:                 "Konflikty!",
		ViewMergeRebaseOptions:              "Widok scalenia/opcje zmiany bazy",
		NotMergingOrRebasing:                "W tej chwili nie scalasz ani nie zmieniasz bazy",
		RecentRepos:                         "Ostatnie repozytoria",
		MergeOptionsTitle:                   "Opcje scalania",
		RebaseOptionsTitle:                  "Opcje zmiany bazy",
		ConflictsResolved:                   "Wszystkie konflikty scalania rozwiązane. Kontynuować?",
		NoRoom:                              "Brak miejsca",
		YouAreHere:                          "JESTEŚ TU",
		RewordNotSupported:                  "Przeredagowanie commitów podczas interaktywnej zmiany bazy nie jest obecnie wspierane",
		CherryPickCopy:                      "Kopiuj commit (przebieranie)",
		CherryPickCopyRange:                 "Kopiuj zakres commitów (przebieranie)",
		PasteCommits:                        "Wklej commity (przebieranie)",
		SureCherryPick:                      "Czy na pewno chcesz przebierać w skopiowanych commitach na tej gałęzi?",
		CherryPick:                          "Przebieranie",
		Donate:                              "Wesprzyj",
		PrevLine:                            "Poprzednia linia",
		NextLine:                            "Następna linia",
		PrevHunk:                            "Poprzedni kawałek",
		NextHunk:                            "Następny kawałek",
		PrevConflict:                        "Poprzedni konflikt",
		NextConflict:                        "Następny konflikt",
		SelectPrevHunk:                      "Wybierz poprzedni kawałek",
		SelectNextHunk:                      "Wybierz następny kawałek",
		ScrollDown:                          "Przewiń w dół",
		ScrollUp:                            "Przewiń w górę",
		AmendCommitTitle:                    "Popraw commit",
		AmendCommitPrompt:                   "Czy na pewno chcesz poprawić ten commit plikami z poczekalni?",
		DeleteCommitTitle:                   "Usuń commit",
		DeleteCommitPrompt:                  "Czy na pewno usunąć ten commit?",
		SquashingStatus:                     "Spłaszczanie",
		FixingStatus:                        "Naprawianie",
		DeletingStatus:                      "Usuwanie",
		MovingStatus:                        "Przenoszenie",
		RebasingStatus:                      "Zmiana bazy",
		AmendingStatus:                      "Poprawianie",
		CherryPickingStatus:                 "Przebieranie",
		CommitFiles:                         "Pliki commita",
		ViewItemFiles:                       "Przeglądaj pliki commita",
		CommitFilesTitle:                    "Pliki commita",
		CheckoutCommitFile:                  "Plik wybierania",
		DiscardOldFileChange:                "Porzuć zmiany commita dla tego pliku",
		DiscardFileChangesTitle:             "Porzuć zmiany w pliku",
		DiscardFileChangesPrompt:            "Czy na pewno porzucić zmiany w tym pliku? Jeśli ten plik był utworzony w tym commicie, zostanie usunięty",
		DisabledForGPG:                      "Funkcja niedostępna dla użytkowników GPG",
		CreateRepo:                          "Nie jesteś w repozytorium. Utworzyć nowe repozytorium Gita? (y/n): ",
		AutoStashTitle:                      "Automatyczny schowek",
		AutoStashPrompt:                     "Musisz schować i wyciągnąć zmiany żeby je przenosić. Robić to automatycznie? (enter/esc)",
		StashPrefix:                         "Automatyczne dodawanie zmian do schowka dla ",
		ViewDiscardOptions:                  "Pokaż opcje porzucania zmian",
		Cancel:                              "Anuluj",
		DiscardAllChanges:                   "Porzuć wszystkie zmiany",
		DiscardUnstagedChanges:              "Porzuć zmiany poza poczekalnią",
		DiscardAllChangesToAllFiles:         "Wytnij drzewo robocze",
		DiscardAnyUnstagedChanges:           "Porzuć zmiany poza poczekalnią",
		DiscardUntrackedFiles:               "Porzuć pliki nieśledzone",
		HardReset:                           "Twardy reset",
		ViewResetOptions:                    "Wyświetl opcje resetu",
		CreateFixupCommitDescription:        "Utwórz commit naprawczy dla tego commita",
		SquashAboveCommits:                  `Spłaszcz wszystkie commity naprawcze powyżej zaznaczonych commitów (autosquash)`,
		SureSquashAboveCommits:              `Na pewno chcesz spłaszczyć wszystkie commity naprawcze powyżej {{.commit}}?`,
		CreateFixupCommit:                   `Utwóż commit naprawczy`,
		SureCreateFixupCommit:               `Na pewno utworzyć commit naprawczy dla commita {{.commit}}?`,
		ExecuteCustomCommand:                "Wykonaj własną komendę",
		CustomCommand:                       "Własna komenda:",
		CommitChangesWithoutHook:            "Zatwierdź zmiany bez skryptu pre-commit",
		SkipHookPrefixNotConfigured:         "Nie masz skonfigurowanego prefixa komunikatu commita dla pomijania skryptów. Ustaw `git.skipHookPrefix = 'WIP'` w konfiguracji",
		ResetTo:                             "Zresetuj do",
		PressEnterToReturn:                  "Wciśnij enter żeby wrócić do lazygit",
		ViewStashOptions:                    "Wyświetl opcje schowka",
		StashAllChanges:                     "Przechowaj zmiany",
		StashAllChangesKeepIndex:            "Przechowaj zmiany z poczekalni",
		StashOptions:                        "Opcje schowka",
		NotARepository:                      "Błąd: nie jesteś w repozytorium",
		Jump:                                "Przeskocz do panelu",
		ExitCustomPatchBuilder:              `Wyście z trybu "linia po linii"`,
		EnterUpstream:                       `Podaj gałąź nadrzędną jako '<remote> <branchname>'`,
		ReturnToRemotesList:                 `Wróć do listy repozytoriów zdalnych`,
		IgnoreTracked:                       "Ignoruj plik śledzony",
		IgnoreTrackedPrompt:                 "Na pewno ignorować plik śledzony?",
		CommitPrefixPatternError:            "Błąd we wzorcu prefixu commita",
		NoFilesStagedTitle:                  "Brak plików w poczekalni",
		NoFilesStagedPrompt:                 "Nie masz plików w poczekalni. Zatwierdzić wszystkie pliki?",
		BranchNotFoundTitle:                 "Nie znaleziono gałęzi",
		BranchNotFoundPrompt:                "Nie znaleziono gałęzi. Utwórz nową gałąź ",
		PullRequestURLCopiedToClipboard:     "URL żądania ściągnięcia skopiowany do schowka",
		CommitMessageCopiedToClipboard:      "Komunikat commita skopiowany do schowka",
		CopiedToClipboard:                   "Skopiowany do schowka",
		CreatePullRequestOptions:            "Utwórz opcje żądania ściągnięcia",
		ConfirmRevertCommit:                 "Czy na pewno chcesz obrócić {{.selectedCommit}}?",
	}
}
