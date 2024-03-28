#ifdef _RUNGAME_
#define _RUNGAME_

#define RUNGAME_C_API __declspec(dllexport)

#ifdef __cplusplus
extern "C"
{
#endif
	//RUNGAME_C_API void __stdcall RunGame(char* path);
	RUNGAME_C_API void __stdcall RunGame(wchar_t* path);
#ifdef __cplusplus
}
#endif

#endif // RUNGAME_H

